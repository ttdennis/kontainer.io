package kmi_test

import (
	"reflect"

	"github.com/ttdennis/kontainer.io/pkg/kmi"
	"github.com/ttdennis/kontainer.io/pkg/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KMI", func() {
	Describe("Create Service", func() {
		It("Should create service", func() {
			kmiService, err := kmi.NewService(testutils.NewMockDB())
			Ω(err).ShouldNot(HaveOccurred())
			Expect(kmiService).ToNot(BeZero())
		})

		It("Should return db error", func() {
			db := testutils.NewMockDB()
			db.SetError(1)
			_, err := kmi.NewService(db)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Extract Functions Error Handling", func() {
		Context("Extract", func() {
			It("Should return an error if src not exists", func() {
				err := kmi.Extract("blub", &kmi.Content{})
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if src is not a tar ball", func() {
				err := kmi.Extract("extract.go", &kmi.Content{})
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("Choose Source", func() {
			It("Should return an error if outsrc is not of type JSON", func() {
				var src interface{}
				src = "test"
				err := kmi.ChooseSource(&src, []byte("a"), reflect.Slice, "")
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if src is not of the right Kind", func() {
				var src interface{}
				src = make(map[string]string)
				err := kmi.ChooseSource(&src, []byte("a"), reflect.Slice, "")
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("Extract String Map", func() {
			It("Should return an error if the source value is not of Kind map", func() {
				err := kmi.ExtractStringMap(reflect.ValueOf(1), nil, nil)
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if a restricted value (int) is in the source", func() {
				restriction := make(map[reflect.Kind]bool)
				restriction[reflect.Int] = true
				m := make(map[string]interface{})
				m["key"] = 1.0

				err := kmi.ExtractStringMap(reflect.ValueOf(m), nil, restriction)
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if a restricted value (string) is in the source", func() {
				restriction := make(map[reflect.Kind]bool)
				restriction[reflect.String] = true
				m := make(map[string]interface{})
				m["key"] = "value"

				err := kmi.ExtractStringMap(reflect.ValueOf(m), nil, restriction)
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if an unimplemented value is in the source", func() {
				m := make(map[string]interface{})
				m["key"] = []byte("value")

				err := kmi.ExtractStringMap(reflect.ValueOf(m), nil, nil)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("Get String Map", func() {
			It("Should return an error if the source value is corrupted", func() {
				err := kmi.GetStringMap([]byte("src"), nil, nil, "name", nil)
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if source map is malformmatted", func() {
				src := make(map[string]interface{})
				r := make(map[reflect.Kind]bool)

				src["key"] = "value"
				r[reflect.String] = true

				err := kmi.GetStringMap(src, nil, nil, "name", r)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("Get String Slice", func() {
			It("Should return an error if the source value is corrupted", func() {
				err := kmi.GetStringSlice(0xDEADBEEF, nil, nil, "name")
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if the source slice is malformatted", func() {
				src := make([]interface{}, 1)
				src[0] = 0xDEADBEEF
				err := kmi.GetStringSlice(src, nil, nil, "name")
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("Get Frontend", func() {
			It("Should return an error if the source value is corrupted", func() {
				err := kmi.GetFrontend(0xDEADBEEF, nil, nil)
				Ω(err).Should(HaveOccurred())
			})

			slice := make([]interface{}, 1)
			m := make(map[string]interface{})
			params := make(map[string]interface{})
			slice[0] = m

			It("Should return an error if the template value is not of type string", func() {
				m["template"] = 2
				err := kmi.GetFrontend(slice, nil, nil)
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if the parameters value is not of type json", func() {
				m["template"] = "template"
				m["parameters"] = "not json"
				err := kmi.GetFrontend(slice, nil, nil)
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if the parameters value is not a valid string map", func() {
				params["test"] = 0xDEADBEEF
				m["parameters"] = params
				err := kmi.GetFrontend(slice, nil, nil)
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if an unexpected property exists in the file", func() {
				params["test"] = "test"
				m["unexpected"] = "property"
				err := kmi.GetFrontend(slice, nil, nil)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("Get Data", func() {
			content := &kmi.Content{}
			It("Should return an error if module json is corrupt", func() {
				err := kmi.GetData(content, &kmi.KMI{})
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if commands json is corrupt", func() {
				content.Module = []byte(`{
          "commands": "",
          "env": "",
          "cmd": "",
          "interfaces": "",
          "imports": "",
          "frontend": ""
          }`)
				err := kmi.GetData(content, &kmi.KMI{})
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if environment json is corrupt", func() {
				content.Cmd = []byte("{}")
				err := kmi.GetData(content, &kmi.KMI{})
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if interfaces json is corrupt", func() {
				content.Env = []byte("{}")
				err := kmi.GetData(content, &kmi.KMI{})
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if frontend json is corrupt", func() {
				content.Interfaces = []byte("{}")
				err := kmi.GetData(content, &kmi.KMI{})
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an error if Imports json is corrupt", func() {
				content.Frontend = []byte("[]")
				err := kmi.GetData(content, &kmi.KMI{})
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Add KMI", func() {
		It("Should Add KMI", func() {
			kmiService, _ := kmi.NewService(testutils.NewMockDB())
			id, err := kmiService.AddKMI("test.kmi")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(id).Should(BeEquivalentTo(1))
		})

		It("Should return an Error if path is broken", func() {
			kmiService, _ := kmi.NewService(testutils.NewMockDB())
			_, err := kmiService.AddKMI("path")
			Ω(err).Should(HaveOccurred())
		})

		It("Should handle DB errors", func() {
			db := testutils.NewMockDB()
			kmiService, _ := kmi.NewService(db)
			db.SetError(2)
			_, err := kmiService.AddKMI("test.kmi")
			Ω(err).Should(HaveOccurred())
		})

		It("Should return an error if module already exists", func() {
			db := testutils.NewMockDB()
			kmiService, _ := kmi.NewService(db)
			kmiService.AddKMI("test.kmi")
			_, err := kmiService.AddKMI("test.kmi")
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Remove KMI", func() {
		db := testutils.NewMockDB()
		kmiService, _ := kmi.NewService(db)
		It("Should Remove KMI", func() {
			id, _ := kmiService.AddKMI("test.kmi")
			err := kmiService.RemoveKMI(id)
			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("Get KMI", func() {
		db := testutils.NewMockDB()
		kmiService, _ := kmi.NewService(db)
		It("Should return the requested kmi", func() {
			id, _ := kmiService.AddKMI("test.kmi")
			k := &kmi.KMI{}
			err := kmiService.GetKMI(id, k)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(k.ID).Should(BeEquivalentTo(id))
		})

		It("Should handle db Errors", func() {
			db.SetError(1)
			err := kmiService.GetKMI(0, &kmi.KMI{})
			Ω(err).Should(HaveOccurred())

			db.SetError(2)
			err = kmiService.GetKMI(0, &kmi.KMI{})
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("KMI", func() {
		db := testutils.NewMockDB()
		kmiService, _ := kmi.NewService(db)
		id, _ := kmiService.AddKMI("test.kmi")
		It("Should return all kmdi", func() {
			res := []kmi.KMDI{}
			err := kmiService.KMI(&res)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res[0].ID).To(BeEquivalentTo(id))
		})

		It("Should handle DB errors", func() {
			db.SetError(1)
			err := kmiService.KMI(&[]kmi.KMDI{})
			Ω(err).Should(HaveOccurred())
		})
	})
})