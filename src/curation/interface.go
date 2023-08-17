package curation

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -lstdc++ -L${SRCDIR}/cpp/lib -lcategorizer -lfilter
#include "cpp/categorizer.h"
#include "cpp/filter.h"
*/
import "C"
import "unsafe"

type Category int

const (
    SCIENCE Category = iota
    TECH
    ART
    NONE
)

func (c Category) String() string {
    switch c {
    case SCIENCE:
        return "SCIENCE"
    case TECH:
        return "TECH"
    case ART:
        return "ART"
    case NONE:
        return "NONE"
    default:
        return "UNKNOWN"
    }
}

func Categorize(description string) []Category {
    desc := C.CString(description)
    defer C.my_free(unsafe.Pointer(desc))

    var size C.int
    categoriesC := C.categorize(desc, &size)
    defer C.free_categories(categoriesC)

    categoriesGo := make([]Category, int(size))

    for i := 0; i < int(size); i++ {
        categoriesGo[i] = Category(*categoriesC)
        categoriesC = (*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(categoriesC)) + unsafe.Sizeof(*categoriesC)))
    }

    return categoriesGo
}

func IsRelevant(description string) bool {
    desc := C.CString(description)
    defer C.my_free(unsafe.Pointer(desc))

    return bool(C.isRelevant(desc))
}