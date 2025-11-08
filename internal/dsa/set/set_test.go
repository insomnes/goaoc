package set

import (
	"fmt"
	"testing"
)

func uniqCount[T comparable](xs []T) int {
	m := make(map[T]struct{}, len(xs))
	for _, x := range xs {
		m[x] = struct{}{}
	}
	return len(m)
}

func requireSetHasAll[T comparable](t *testing.T, s *Set[T], want []T) {
	t.Helper()
	if s.Size() != uniqCount(want) {
		t.Fatalf("size mismatch: got %d, want %d", s.Size(), uniqCount(want))
	}
	for _, w := range want {
		if !s.Contains(w) {
			t.Fatalf("missing element: %v", w)
		}
	}
}

func requireElementsMatch[T comparable](t *testing.T, s *Set[T]) {
	t.Helper()
	got := s.Elements()
	seen := make(map[T]struct{}, len(got))
	for _, g := range got {
		seen[g] = struct{}{}
		if !s.Contains(g) {
			t.Fatalf("Elements() returned value not in set: %v", g)
		}
	}
	if len(seen) != s.Size() {
		t.Fatalf("Elements() count mismatch: got %d unique, want %d", len(seen), s.Size())
	}
}

func TestNewSet_Empty(t *testing.T) {
	s := NewSet[int]()
	if !s.IsEmpty() || s.Size() != 0 {
		t.Fatalf("new set not empty: IsEmpty=%v Size=%d", s.IsEmpty(), s.Size())
	}
}

func TestNewSetCap(t *testing.T) {
	s := NewSetCap[int](10)
	if !s.IsEmpty() || s.Size() != 0 {
		t.Fatalf("new set not empty: IsEmpty=%v Size=%d", s.IsEmpty(), s.Size())
	}
}

func TestNewSetCap_Negative(t *testing.T) {
	s := NewSetCap[int](-5)
	if !s.IsEmpty() || s.Size() != 0 {
		t.Fatalf("new set not empty: IsEmpty=%v Size=%d", s.IsEmpty(), s.Size())
	}
}

func TestCopyEmptySet(t *testing.T) {
	s1 := NewSet[int]()
	s2 := s1.Copy()
	if !s2.IsEmpty() || s2.Size() != 0 {
		t.Fatalf("copied set not empty: IsEmpty=%v Size=%d", s2.IsEmpty(), s2.Size())
	}
}

func TestFromString(t *testing.T) {
	s := FromString("hello")
	requireSetHasAll(t, s, []rune{'h', 'e', 'l', 'o'})
}

func TestFromEmptyString(t *testing.T) {
	s := FromString("")
	if !s.IsEmpty() || s.Size() != 0 {
		t.Fatalf("FromString(\"\") not empty: IsEmpty=%v Size=%d", s.IsEmpty(), s.Size())
	}
}

func TestFromString_Uniqueness(t *testing.T) {
	s := FromString("aaabbbccc")
	requireSetHasAll(t, s, []rune{'a', 'b', 'c'})
}

func TestFromString_Unicode(t *testing.T) {
	s := FromString("こんにちはこんにちは")
	requireSetHasAll(t, s, []rune{'こ', 'ん', 'に', 'ち', 'は'})
}

func TestFromString_Emoji(t *testing.T) {
	s := FromString("😀😃😄😁😀😃")
	requireSetHasAll(t, s, []rune{'😀', '😃', '😄', '😁'})
}

func TestAddContainsRemove(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(1) // duplicate
	if !s.Contains(1) || s.Size() != 1 {
		t.Fatalf("add/contains failed: contains=%v size=%d", s.Contains(1), s.Size())
	}
	s.Add(2)
	requireSetHasAll(t, s, []int{1, 2})

	s.Remove(1)
	if s.Contains(1) || s.Size() != 1 {
		t.Fatalf("remove failed: contains=%v size=%d", s.Contains(1), s.Size())
	}
	s.Remove(42) // no-op
	requireSetHasAll(t, s, []int{2})
}

func TestClearAndIsEmpty(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	s.Clear()
	if !s.IsEmpty() || s.Size() != 0 {
		t.Fatalf("clear failed: IsEmpty=%v Size=%d", s.IsEmpty(), s.Size())
	}
}

func TestElements(t *testing.T) {
	s := FromSlice([]int{1, 2, 2, 3})
	requireElementsMatch(t, s)
}

func TestCopyIndependence(t *testing.T) {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := s1.Copy()
	s1.Add(4)
	s2.Remove(2)

	if s1.Contains(2) == false {
		t.Fatalf("modifying copy affected original")
	}
	if s2.Contains(4) {
		t.Fatalf("modifying original affected copy")
	}
	requireSetHasAll(t, s1, []int{1, 2, 3, 4})
	requireSetHasAll(t, s2, []int{1, 3})
}

func TestFromSlice(t *testing.T) {
	s := FromSlice([]int{1, 1, 2, 3, 3})
	requireSetHasAll(t, s, []int{1, 2, 3})
}

func TestFromIter(t *testing.T) {
	iter := func(yield func(int) bool) {
		for _, v := range []int{1, 2, 2, 3} {
			if !yield(v) {
				return
			}
		}
	}
	s := FromIter(iter)
	requireSetHasAll(t, s, []int{1, 2, 3})
}

func TestIter_EarlyStop(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 4, 5})
	iter := s.Iter()

	count := 0
	iter(func(_ int) bool {
		count++
		return count < 3 // stop after seeing 3rd value
	})
	if count != 3 {
		t.Fatalf("early stop failed: iterated %d elements, want 3", count)
	}
}

func TestUnion(t *testing.T) {
	a := FromSlice([]int{1, 2, 3})
	b := FromSlice([]int{3, 4})
	u := a.Union(b)
	requireSetHasAll(t, u, []int{1, 2, 3, 4})
	// original sets intact
	requireSetHasAll(t, a, []int{1, 2, 3})
	requireSetHasAll(t, b, []int{3, 4})
}

func TestIntersection(t *testing.T) {
	a := FromSlice([]int{1, 2, 3})
	b := FromSlice([]int{3, 4})
	i := a.Intersection(b)
	requireSetHasAll(t, i, []int{3})

	empty := a.Intersection(FromSlice([]int{5, 6}))
	if !empty.IsEmpty() {
		t.Fatalf("expected empty intersection")
	}
}

func TestDifference(t *testing.T) {
	a := FromSlice([]int{1, 2, 3, 5})
	b := FromSlice([]int{3, 4})
	d1 := a.Difference(b) // elements in a not in b
	requireSetHasAll(t, d1, []int{1, 2, 5})

	d2 := b.Difference(a) // elements in b not in a
	requireSetHasAll(t, d2, []int{4})
}

func TestSubsetSuperset(t *testing.T) {
	a := FromSlice([]int{1, 2})
	b := FromSlice([]int{1, 2, 3})

	if !a.IsSubset(b) {
		t.Fatalf("a should be subset of b")
	}
	if !b.IsSuperset(a) {
		t.Fatalf("b should be superset of a")
	}
	if b.IsSubset(a) {
		t.Fatalf("b should not be subset of a")
	}
}

func TestEqual(t *testing.T) {
	a := FromSlice([]int{1, 2, 3})
	b := FromSlice([]int{3, 2, 1})
	c := FromSlice([]int{1, 2})
	if !a.Equal(b) {
		t.Fatalf("sets with same elements should be equal")
	}
	if a.Equal(c) {
		t.Fatalf("sets with different elements should not be equal")
	}
}

func TestClearIf(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 4, 5, 6})
	s.ClearIf(func(x int) bool { return x%2 == 0 }) // remove evens
	requireSetHasAll(t, s, []int{1, 3, 5})
}

func TestRetainIf(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 4, 5})
	s.RetainIf(func(x int) bool { return x > 2 }) // keep > 2
	requireSetHasAll(t, s, []int{3, 4, 5})
}

func TestGenericStrings(t *testing.T) {
	s := FromSlice([]string{"a", "b", "b"})
	if s.Size() != 2 || !s.Contains("a") || !s.Contains("b") {
		t.Fatalf("string set basic ops failed")
	}
	u := s.Union(FromSlice([]string{"b", "c"}))
	requireSetHasAll(t, u, []string{"a", "b", "c"})
}

func TestUnionWithEmpty(t *testing.T) {
	a := FromSlice([]int{1, 2, 3})
	empty := NewSet[int]()

	u1 := a.Union(empty)
	u2 := empty.Union(a)

	if !u1.Equal(a) {
		t.Fatalf("union with empty should equal original: %v", u1.Elements())
	}
	if !u2.Equal(a) {
		t.Fatalf("empty union original should equal original: %v", u2.Elements())
	}
}

func TestIntersectionWithSelf(t *testing.T) {
	a := FromSlice([]string{"x", "y", "z"})
	i := a.Intersection(a)
	if !i.Equal(a) {
		t.Fatalf("intersection with self should equal self: %v", i.Elements())
	}
}

func TestDifferenceProperties(t *testing.T) {
	a := FromSlice([]int{1, 2, 3, 4})
	b := FromSlice([]int{3, 4, 5})

	ab := a.Difference(b) // 1,2
	ba := b.Difference(a) // 5

	for _, v := range []int{1, 2} {
		if !ab.Contains(v) {
			t.Fatalf("a\\b missing %d", v)
		}
	}
	if ab.Contains(3) || ab.Contains(4) || ab.Contains(5) {
		t.Fatalf("a\\b contains undesired elements: %v", ab.Elements())
	}
	if !ba.Equal(FromSlice([]int{5})) {
		t.Fatalf("b\\a should be {5}: %v", ba.Elements())
	}

	// a \ a = ∅, a \ ∅ = a
	if !a.Difference(a).IsEmpty() {
		t.Fatalf("a\\a should be empty")
	}
	if !a.Difference(NewSet[int]()).Equal(a) {
		t.Fatalf("a\\∅ should be a")
	}
}

func TestSubsetSupersetEqualEmpty(t *testing.T) {
	empty1 := NewSet[int]()
	empty2 := NewSet[int]()
	if !empty1.Equal(empty2) {
		t.Fatalf("two empty sets should be equal")
	}
	if !empty1.IsSubset(empty2) || !empty1.IsSuperset(empty2) {
		t.Fatalf("empty should be subset and superset of empty")
	}

	a := FromSlice([]int{1, 2})
	b := FromSlice([]int{1, 2})
	if !a.IsSubset(b) || !a.IsSuperset(b) || !a.Equal(b) {
		t.Fatalf("equal sets should be subset/superset/equal")
	}
}

func TestIterStopsWhenYieldFalse(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 4, 5})
	iter := s.Iter()

	s2 := NewSet[int]()

	iter(func(v int) bool {
		s2.Add(v)
		if s2.Size() >= 2 {
			return false // stop after 2 elements
		}
		return true
	})
	if s2.Size() > 2 {
		t.Fatalf("expected early stop after 2 element, got %d", s2.Size())
	}
}

func TestIterDeleteDuringIteration(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 4})
	iter := s.Iter()

	visited := make(map[int]struct{})
	iter(func(v int) bool {
		visited[v] = struct{}{}
		if v%2 == 0 {
			s.Remove(v) // delete even during iteration
		}
		return true
	})

	// Set should now only contain odds.
	for _, want := range []int{1, 3} {
		if !s.Contains(want) {
			t.Fatalf("expected to retain %d", want)
		}
	}
	for _, bad := range []int{2, 4} {
		if s.Contains(bad) {
			t.Fatalf("expected to remove %d", bad)
		}
	}
	// Ensure we actually visited at least one even and one odd.
	_, visited2 := visited[2]
	_, visited4 := visited[4]
	if !visited2 && !visited4 {
		t.Fatalf("expected to visit an even element at least once, visited=%v", visited)
	}

	_, visited1 := visited[1]
	_, visited3 := visited[3]

	if !visited1 && !visited3 {
		t.Fatalf("expected to visit an odd element at least once, visited=%v", visited)
	}
}

func TestRetainIf_NoChange_And_AllRemove(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	s.RetainIf(func(x int) bool { return x >= 1 }) // no change
	if s.Size() != 3 {
		t.Fatalf("retain-if no-op changed size: %d", s.Size())
	}

	s.RetainIf(func(int) bool { return false }) // remove all
	if !s.IsEmpty() {
		t.Fatalf("retain-if remove all did not empty set")
	}
}

func TestClearIf_RemoveAllOrNone(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	s.ClearIf(func(int) bool { return false }) // none removed
	if s.Size() != 3 {
		t.Fatalf("clear-if false removed elements")
	}

	s.ClearIf(func(int) bool { return true }) // all removed
	if !s.IsEmpty() {
		t.Fatalf("clear-if true did not remove all")
	}
}

func TestGenericStructKeys(t *testing.T) {
	type key struct {
		A int
		B string
	}
	a := key{1, "x"}
	b := key{1, "y"}
	c := key{2, "x"}

	s := NewSet[key]()
	s.Add(a)
	s.Add(b)
	s.Add(c)
	if s.Size() != 3 {
		t.Fatalf("unexpected size: %d", s.Size())
	}
	if !s.Contains(key{1, "x"}) || !s.Contains(key{2, "x"}) {
		t.Fatalf("missing expected keys")
	}
}

func TestElementsUniqueness(t *testing.T) {
	s := FromSlice([]int{1, 1, 2, 3, 3, 3})
	elems := s.Elements()

	seen := make(map[int]int)
	for _, e := range elems {
		seen[e]++
	}
	for k, n := range seen {
		if n > 1 {
			t.Fatalf("duplicate element from Elements: %d seen %d times", k, n)
		}
		if !s.Contains(k) {
			t.Fatalf("Elements returned value not in set: %d", k)
		}
	}
	if len(seen) != s.Size() {
		t.Fatalf("Elements unique count mismatch: got %d want %d", len(seen), s.Size())
	}
}

func TestFromIter_CustomIterator(t *testing.T) {
	// iterator that can short-circuit based on yield return value
	source := []int{10, 20, 30, 40}
	iter := func(yield func(int) bool) {
		for _, v := range source {
			if !yield(v) {
				return
			}
		}
	}

	s := FromIter(iter)
	for _, v := range source {
		if !s.Contains(v) {
			t.Fatalf("FromIter missing %d", v)
		}
	}

	// sanity: ensure our iterator can stop early when used directly
	resultSet := NewSet[int]()
	iter(func(v int) bool {
		resultSet.Add(v)
		if resultSet.Size() >= 1 {
			return false // stop after first
		}
		return true
	})
	if resultSet.Size() != 1 {
		t.Fatalf("iterator did not stop early: count=%d", resultSet.Size())
	}
}

func TestStringerElementsDebug(t *testing.T) {
	s := FromSlice([]int{7, 8, 9})
	_ = fmt.Sprint(s.Elements()) // ensure no panic when formatting
}

func containsSubstring(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr || len(s) > len(substr) && (s[0:len(substr)] == substr || containsSubstring(s[1:], substr)))
}

func TestString(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	str := s.String()
	expectedSubstrings := []string{"1", "2", "3"}
	for _, substr := range expectedSubstrings {
		if !containsSubstring(str, substr) {
			t.Fatalf("String() output missing %q: got %q", substr, str)
		}
	}
	if str[0] != '{' || str[len(str)-1] != '}' {
		t.Fatalf("String() output not properly formatted: got %q", str)
	}

	commaCount := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ',' {
			commaCount++
		}
	}
	expectedCommas := s.Size() - 1
	if commaCount != expectedCommas {
		t.Fatalf(
			"String() output has incorrect number of elements: got %d commas, want %d",
			commaCount,
			expectedCommas,
		)
	}

	seen := make(map[string]struct{})
	current := ""
	for i := 1; i < len(str)-1; i++ { // skip '{' and '}'
		if str[i] == ',' {
			trimmed := current
			seen[trimmed] = struct{}{}
			current = ""
			i++ // skip space after comma
		} else {
			current += string(str[i])
		}
	}
	if current != "" {
		trimmed := current
		seen[trimmed] = struct{}{}
	}
	if len(seen) != s.Size() {
		t.Fatalf("String() output has duplicate elements: %q", str)
	}
}

func TestString_EmptySet(t *testing.T) {
	s := NewSet[int]()
	str := s.String()
	if str != "{}" {
		t.Fatalf("String() of empty set should be {}, got %q", str)
	}
}

func TestAddAllRemoveAllHasAllAny(t *testing.T) {
	s := NewSet[int]()
	s.AddAll(1, 2, 2, 3)
	requireSetHasAll(t, s, []int{1, 2, 3})

	if !s.HasAll(1, 2) {
		t.Fatalf("HasAll should be true")
	}
	if s.HasAll(1, 4) {
		t.Fatalf("HasAll should be false")
	}
	if !s.HasAny(0, 3) {
		t.Fatalf("HasAny should be true")
	}
	if s.HasAny(9, 8) {
		t.Fatalf("HasAny should be false")
	}

	s.RemoveAll(2, 9) // 9 is no-op
	requireSetHasAll(t, s, []int{1, 3})
}

func TestUpdateIntersectWithDifferenceWith(t *testing.T) {
	a := FromSlice([]int{1, 2})
	b := FromSlice([]int{2, 3})

	c := a.Copy()
	c.Update(b) // {1,2,3}
	requireSetHasAll(t, c, []int{1, 2, 3})
	// originals intact
	requireSetHasAll(t, a, []int{1, 2})
	requireSetHasAll(t, b, []int{2, 3})

	d := FromSlice([]int{1, 2, 3, 4})
	d.IntersectWith(b) // keep {2,3}
	requireSetHasAll(t, d, []int{2, 3})

	e := FromSlice([]int{1, 2, 3})
	e.DifferenceWith(b) // remove 2,3 -> {1}
	requireSetHasAll(t, e, []int{1})
}

func TestIsDisjointOverlaps(t *testing.T) {
	a := FromSlice([]int{1, 2})
	b := FromSlice([]int{3, 4})
	c := FromSlice([]int{2, 3})

	if !a.IsDisjoint(b) {
		t.Fatalf("expected disjoint")
	}
	if a.Overlaps(b) {
		t.Fatalf("unexpected overlap")
	}
	if a.IsDisjoint(c) {
		t.Fatalf("should not be disjoint")
	}
	if !a.Overlaps(c) {
		t.Fatalf("should overlap")
	}
}

func TestSymmetricDifferenceProperties(t *testing.T) {
	a := FromSlice([]int{1, 2, 3})
	b := FromSlice([]int{3, 4, 5})

	sd := a.SymmetricDifference(b) // {1,2,4,5}
	requireSetHasAll(t, sd, []int{1, 2, 4, 5})

	// Commutative
	sd2 := b.SymmetricDifference(a)
	if !sd.Equal(sd2) {
		t.Fatalf("symmetric difference should be commutative")
	}

	// With self -> empty
	if !a.SymmetricDifference(a).IsEmpty() {
		t.Fatalf("A △ A should be empty")
	}
}

func TestPop(t *testing.T) {
	empty := NewSet[int]()
	if _, ok := empty.Pop(); ok {
		t.Fatalf("Pop on empty should return ok=false")
	}

	s := FromSlice([]int{10, 20, 30})
	before := s.Size()
	val, ok := s.Pop()
	if !ok {
		t.Fatalf("expected ok=true")
	}
	if s.Size() != before-1 {
		t.Fatalf("size should decrease by 1, got %d want %d", s.Size(), before-1)
	}
	if s.Contains(val) {
		t.Fatalf("popped value should be removed")
	}
	// remaining should be a subset of original
	for _, v := range []int{10, 20, 30} {
		if v != val && !s.Contains(v) && s.Size() != 1 { // after one pop we still have 2 elements
			t.Fatalf("unexpected missing element after pop: %d", v)
		}
	}
}

func TestFilter(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 4, 5, 6})
	evens := s.Filter(func(x int) bool { return x%2 == 0 })
	requireSetHasAll(t, evens, []int{2, 4, 6})

	none := s.Filter(func(int) bool { return false })
	if !none.IsEmpty() {
		t.Fatalf("filter false should be empty")
	}
	all := s.Filter(func(int) bool { return true })
	if !all.Equal(s) {
		t.Fatalf("filter true should equal original")
	}
}

func TestMap(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	mapped := Map[int, string](s, func(x int) string { return string(rune('a' + x - 1)) })
	requireSetHasAll(t, mapped, []string{"a", "b", "c"})

	// mapping to non-unique values should collapse to set uniqueness
	s2 := FromSlice([]int{1, 2, 3, 4})
	parity := Map[int, int](s2, func(x int) int { return x % 2 })
	// expect {0,1}
	if parity.Size() != 2 || !parity.HasAll(0, 1) {
		t.Fatalf("map should produce unique collapsed set, got %v", parity.Elements())
	}
}
