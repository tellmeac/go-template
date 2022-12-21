// Code generated by ent, DO NOT EDIT.

package templates

import (
	"entgo.io/ent/dialect/sql"
	"github.com/tellmeac/go-template/internal/infrastructure/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Template applies equality check predicate on the "template" field. It's identical to TemplateEQ.
func Template(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemplate), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Templates {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Templates {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TemplateEQ applies the EQ predicate on the "template" field.
func TemplateEQ(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemplate), v))
	})
}

// TemplateNEQ applies the NEQ predicate on the "template" field.
func TemplateNEQ(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTemplate), v))
	})
}

// TemplateIn applies the In predicate on the "template" field.
func TemplateIn(vs ...string) predicate.Templates {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTemplate), v...))
	})
}

// TemplateNotIn applies the NotIn predicate on the "template" field.
func TemplateNotIn(vs ...string) predicate.Templates {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTemplate), v...))
	})
}

// TemplateGT applies the GT predicate on the "template" field.
func TemplateGT(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTemplate), v))
	})
}

// TemplateGTE applies the GTE predicate on the "template" field.
func TemplateGTE(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTemplate), v))
	})
}

// TemplateLT applies the LT predicate on the "template" field.
func TemplateLT(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTemplate), v))
	})
}

// TemplateLTE applies the LTE predicate on the "template" field.
func TemplateLTE(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTemplate), v))
	})
}

// TemplateContains applies the Contains predicate on the "template" field.
func TemplateContains(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTemplate), v))
	})
}

// TemplateHasPrefix applies the HasPrefix predicate on the "template" field.
func TemplateHasPrefix(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTemplate), v))
	})
}

// TemplateHasSuffix applies the HasSuffix predicate on the "template" field.
func TemplateHasSuffix(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTemplate), v))
	})
}

// TemplateEqualFold applies the EqualFold predicate on the "template" field.
func TemplateEqualFold(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTemplate), v))
	})
}

// TemplateContainsFold applies the ContainsFold predicate on the "template" field.
func TemplateContainsFold(v string) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTemplate), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Templates) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Templates) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Templates) predicate.Templates {
	return predicate.Templates(func(s *sql.Selector) {
		p(s.Not())
	})
}
