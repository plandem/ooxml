package ooxml

import (
	"fmt"
	"github.com/plandem/ooxml/ml"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//Relationships is a higher level object that wraps OOXML relationships with functionality
type Relationships struct {
	ml   ml.Relationships
	pkg  *PackageInfo
	file *PackageFile
}

//NewRelationships creates and returns relationships
func NewRelationships(f interface{}, pkg *PackageInfo) *Relationships {
	rels := &Relationships{
		pkg: pkg,
	}

	rels.file = NewPackageFile(pkg, f, &rels.ml, nil)
	rels.file.LoadIfRequired(nil)

	if rels.file.IsNew() {
		pkg.ContentTypes().RegisterContent(rels.file.FileName(), "application/vnd.openxmlformats-package.relationships+xml")
		rels.file.MarkAsUpdated()
	}

	return rels
}

//FileName returns file name of relations
func (rels *Relationships) FileName() string {
	return rels.file.fileName
}

//Total returns total number of relationships
func (rels *Relationships) Total() int {
	return len(rels.ml.Relationships)
}

// GetTargetById returns target of relation for provided id
func (rels *Relationships) GetTargetById(id string) string {
	for _, r := range rels.ml.Relationships {
		if r.ID == id {
			return string(r.Target)
		}
	}

	return ""
}

// GetIdByTarget returns id of relation for provided target
func (rels *Relationships) GetIdByTarget(target string) ml.RID {
	for _, r := range rels.ml.Relationships {
		rTarget := r.Target

		switch r.TargetMode {
		case ml.TargetModeInternal:
			//is weird case when link is related?
			if rTarget[0] != '/' {
				rTarget = "/xl/" + rTarget
			}

			if strings.Contains(string(rTarget), target) {
				return ml.RID(r.ID)
			}
		case ml.TargetModeExternal:
			if rTarget == target {
				return ml.RID(r.ID)
			}
		}
	}

	return ""
}

//AddLink adds a new relation of type t to external target - e.g.: url
func (rels *Relationships) AddLink(t ml.RelationType, target string) (int, ml.RID) {
	return rels.add(t, target, ml.TargetModeExternal)
}

//AddFile adds a new relation of type t to internal target, i.e. file inside of package. For simplicity - use absolute paths.
func (rels *Relationships) AddFile(t ml.RelationType, target string) (int, ml.RID) {
	if target[0] != '/' {
		target = "/" + target
	}

	return rels.add(t, target, ml.TargetModeInternal)
}

func (rels *Relationships) add(t ml.RelationType, target string, mode ml.TargetMode) (int, ml.RID) {
	var id int

	regID := regexp.MustCompile(`[\d]+`)
	for _, rel := range rels.ml.Relationships {
		if rid := regID.FindString(rel.ID); rid != "" {
			if rid, err := strconv.ParseUint(rid, 10, 32); err == nil {
				id = int(math.Max(float64(rid), float64(id)))
			}
		}
	}

	rid := fmt.Sprintf("rId%d", id+1)
	relation := ml.Relation{
		ID:         rid,
		Type:       t,
		Target:     target,
		TargetMode: mode,
	}

	rels.ml.Relationships = append(rels.ml.Relationships, relation)
	rels.file.MarkAsUpdated()
	return id, ml.RID(rid)
}

//Remove removes relation with provided rid
func (rels *Relationships) Remove(rid ml.RID) {
	//remove relation
	for i, r := range rels.ml.Relationships {
		if ml.RID(r.ID) == rid {
			rels.ml.Relationships = append(rels.ml.Relationships[:i], rels.ml.Relationships[i+1:]...)
			break
		}
	}

	rels.file.MarkAsUpdated()
}
