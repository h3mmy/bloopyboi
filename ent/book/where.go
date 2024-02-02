// Code generated by ent, DO NOT EDIT.

package book

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDescription, v))
}

// GoodreadsID applies equality check predicate on the "goodreads_id" field. It's identical to GoodreadsIDEQ.
func GoodreadsID(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldGoodreadsID, v))
}

// GoogleVolumeID applies equality check predicate on the "google_volume_id" field. It's identical to GoogleVolumeIDEQ.
func GoogleVolumeID(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldGoogleVolumeID, v))
}

// Isbn10 applies equality check predicate on the "isbn_10" field. It's identical to Isbn10EQ.
func Isbn10(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldIsbn10, v))
}

// Isbn13 applies equality check predicate on the "isbn_13" field. It's identical to Isbn13EQ.
func Isbn13(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldIsbn13, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldDescription, v))
}

// GoodreadsIDEQ applies the EQ predicate on the "goodreads_id" field.
func GoodreadsIDEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldGoodreadsID, v))
}

// GoodreadsIDNEQ applies the NEQ predicate on the "goodreads_id" field.
func GoodreadsIDNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldGoodreadsID, v))
}

// GoodreadsIDIn applies the In predicate on the "goodreads_id" field.
func GoodreadsIDIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldGoodreadsID, vs...))
}

// GoodreadsIDNotIn applies the NotIn predicate on the "goodreads_id" field.
func GoodreadsIDNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldGoodreadsID, vs...))
}

// GoodreadsIDGT applies the GT predicate on the "goodreads_id" field.
func GoodreadsIDGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldGoodreadsID, v))
}

// GoodreadsIDGTE applies the GTE predicate on the "goodreads_id" field.
func GoodreadsIDGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldGoodreadsID, v))
}

// GoodreadsIDLT applies the LT predicate on the "goodreads_id" field.
func GoodreadsIDLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldGoodreadsID, v))
}

// GoodreadsIDLTE applies the LTE predicate on the "goodreads_id" field.
func GoodreadsIDLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldGoodreadsID, v))
}

// GoodreadsIDContains applies the Contains predicate on the "goodreads_id" field.
func GoodreadsIDContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldGoodreadsID, v))
}

// GoodreadsIDHasPrefix applies the HasPrefix predicate on the "goodreads_id" field.
func GoodreadsIDHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldGoodreadsID, v))
}

// GoodreadsIDHasSuffix applies the HasSuffix predicate on the "goodreads_id" field.
func GoodreadsIDHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldGoodreadsID, v))
}

// GoodreadsIDEqualFold applies the EqualFold predicate on the "goodreads_id" field.
func GoodreadsIDEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldGoodreadsID, v))
}

// GoodreadsIDContainsFold applies the ContainsFold predicate on the "goodreads_id" field.
func GoodreadsIDContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldGoodreadsID, v))
}

// GoogleVolumeIDEQ applies the EQ predicate on the "google_volume_id" field.
func GoogleVolumeIDEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDNEQ applies the NEQ predicate on the "google_volume_id" field.
func GoogleVolumeIDNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDIn applies the In predicate on the "google_volume_id" field.
func GoogleVolumeIDIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldGoogleVolumeID, vs...))
}

// GoogleVolumeIDNotIn applies the NotIn predicate on the "google_volume_id" field.
func GoogleVolumeIDNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldGoogleVolumeID, vs...))
}

// GoogleVolumeIDGT applies the GT predicate on the "google_volume_id" field.
func GoogleVolumeIDGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDGTE applies the GTE predicate on the "google_volume_id" field.
func GoogleVolumeIDGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDLT applies the LT predicate on the "google_volume_id" field.
func GoogleVolumeIDLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDLTE applies the LTE predicate on the "google_volume_id" field.
func GoogleVolumeIDLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDContains applies the Contains predicate on the "google_volume_id" field.
func GoogleVolumeIDContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDHasPrefix applies the HasPrefix predicate on the "google_volume_id" field.
func GoogleVolumeIDHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDHasSuffix applies the HasSuffix predicate on the "google_volume_id" field.
func GoogleVolumeIDHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDEqualFold applies the EqualFold predicate on the "google_volume_id" field.
func GoogleVolumeIDEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldGoogleVolumeID, v))
}

// GoogleVolumeIDContainsFold applies the ContainsFold predicate on the "google_volume_id" field.
func GoogleVolumeIDContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldGoogleVolumeID, v))
}

// Isbn10EQ applies the EQ predicate on the "isbn_10" field.
func Isbn10EQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldIsbn10, v))
}

// Isbn10NEQ applies the NEQ predicate on the "isbn_10" field.
func Isbn10NEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldIsbn10, v))
}

// Isbn10In applies the In predicate on the "isbn_10" field.
func Isbn10In(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldIsbn10, vs...))
}

// Isbn10NotIn applies the NotIn predicate on the "isbn_10" field.
func Isbn10NotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldIsbn10, vs...))
}

// Isbn10GT applies the GT predicate on the "isbn_10" field.
func Isbn10GT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldIsbn10, v))
}

// Isbn10GTE applies the GTE predicate on the "isbn_10" field.
func Isbn10GTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldIsbn10, v))
}

// Isbn10LT applies the LT predicate on the "isbn_10" field.
func Isbn10LT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldIsbn10, v))
}

// Isbn10LTE applies the LTE predicate on the "isbn_10" field.
func Isbn10LTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldIsbn10, v))
}

// Isbn10Contains applies the Contains predicate on the "isbn_10" field.
func Isbn10Contains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldIsbn10, v))
}

// Isbn10HasPrefix applies the HasPrefix predicate on the "isbn_10" field.
func Isbn10HasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldIsbn10, v))
}

// Isbn10HasSuffix applies the HasSuffix predicate on the "isbn_10" field.
func Isbn10HasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldIsbn10, v))
}

// Isbn10EqualFold applies the EqualFold predicate on the "isbn_10" field.
func Isbn10EqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldIsbn10, v))
}

// Isbn10ContainsFold applies the ContainsFold predicate on the "isbn_10" field.
func Isbn10ContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldIsbn10, v))
}

// Isbn13EQ applies the EQ predicate on the "isbn_13" field.
func Isbn13EQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldIsbn13, v))
}

// Isbn13NEQ applies the NEQ predicate on the "isbn_13" field.
func Isbn13NEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldIsbn13, v))
}

// Isbn13In applies the In predicate on the "isbn_13" field.
func Isbn13In(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldIsbn13, vs...))
}

// Isbn13NotIn applies the NotIn predicate on the "isbn_13" field.
func Isbn13NotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldIsbn13, vs...))
}

// Isbn13GT applies the GT predicate on the "isbn_13" field.
func Isbn13GT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldIsbn13, v))
}

// Isbn13GTE applies the GTE predicate on the "isbn_13" field.
func Isbn13GTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldIsbn13, v))
}

// Isbn13LT applies the LT predicate on the "isbn_13" field.
func Isbn13LT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldIsbn13, v))
}

// Isbn13LTE applies the LTE predicate on the "isbn_13" field.
func Isbn13LTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldIsbn13, v))
}

// Isbn13Contains applies the Contains predicate on the "isbn_13" field.
func Isbn13Contains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldIsbn13, v))
}

// Isbn13HasPrefix applies the HasPrefix predicate on the "isbn_13" field.
func Isbn13HasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldIsbn13, v))
}

// Isbn13HasSuffix applies the HasSuffix predicate on the "isbn_13" field.
func Isbn13HasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldIsbn13, v))
}

// Isbn13EqualFold applies the EqualFold predicate on the "isbn_13" field.
func Isbn13EqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldIsbn13, v))
}

// Isbn13ContainsFold applies the ContainsFold predicate on the "isbn_13" field.
func Isbn13ContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldIsbn13, v))
}

// HasBookAuthor applies the HasEdge predicate on the "book_author" edge.
func HasBookAuthor() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, BookAuthorTable, BookAuthorPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookAuthorWith applies the HasEdge predicate on the "book_author" edge with a given conditions (other predicates).
func HasBookAuthorWith(preds ...predicate.BookAuthor) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newBookAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
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
func Not(p predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		p(s.Not())
	})
}
