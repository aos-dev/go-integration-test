package tests

import (
	"github.com/beyondstorage/go-storage/v4/types"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLinker(t *testing.T, store types.Storager) {
	Convey("Given a basic Storager", t, func() {
		l, ok := store.(types.Linker)
		So(ok, ShouldBeTrue)

		Convey("When create a link object", func() {
			path := uuid.New().String()
			_ = store.Create(path)

			defer func() {
				err := store.Delete(path)
				if err != nil {
					t.Error(err)
				}
			}()

			target := uuid.New().String()
			o, err := l.CreateLink(path, target)

			defer func() {
				err = store.Delete(target)
				if err != nil {
					t.Error(err)
				}
			}()

			Convey("The first returned error should be nil", func() {
				So(err, ShouldBeNil)
			})

			o, err = l.CreateLink(path, target)

			Convey("The second returned error also should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The object must should be link", func() {
				// Link object's mode must be link.
				So(o.Mode.IsLink(), ShouldBeTrue)
				// Link object's mode should not be read.
				So(o.Mode.IsRead(), ShouldBeFalse)
			})

			Convey("The linkTarget of the object must be the same as the target", func() {
				// The linkTarget must be the same as the target.
				linkTarget, ok := o.GetLinkTarget()

				So(ok, ShouldBeTrue)
				So(linkTarget, ShouldEqual, target)
			})
		})

		Convey("When create a link object to a not existing target", func() {
			path := uuid.New().String()
			_ = store.Create(path)

			defer func() {
				err := store.Delete(path)
				if err != nil {
					t.Error(err)
				}
			}()

			target := uuid.New().String()

			defer func() {
				err := store.Delete(target)
				if err != nil {
					t.Error(err)
				}
			}()

			o, err := l.CreateLink(path, target)

			Convey("The error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			o, err = l.CreateLink(path, target)

			Convey("The second returned error also should not be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The object should be nil", func() {
				So(o, ShouldBeNil)
			})
		})

		Convey("When create a link object from a not existing path", func() {
			path := uuid.New().String()

			defer func() {
				err := store.Delete(path)
				if err != nil {
					t.Error(err)
				}
			}()

			target := uuid.New().String()
			o, err := l.CreateLink(path, target)

			defer func() {
				err = store.Delete(target)
				if err != nil {
					t.Error(err)
				}
			}()

			Convey("The first returned error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			o, err = l.CreateLink(path, target)

			Convey("The second returned error also should not be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The object should be nil", func() {
				So(o, ShouldBeNil)
			})
		})
	})
}
