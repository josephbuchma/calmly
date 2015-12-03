package calmly

import (
	"testing"
)

type TestPanic struct{}
type TestPanic1 struct{ e string }

func (tp TestPanic1) Error() string {
	return tp.e
}

type TestPanic2 struct{}

func TestCalmly(t *testing.T) {
	var catched, finalized bool

	defer func() {
		if r := recover(); r != nil {
			t.Error("Should not catch panic")
		}
	}()

	Try(func() {
		panic(TestPanic1{"PANIC"})
	}).Catch(TestPanic{}, func(e E) {
		t.Error("This should be skipped")
	}).Catch(TestPanic1{}, func(e E) {
		if e.(error).Error() != "PANIC" {
			t.Errorf("Expected PANIC, got %s", e.(error).Error())
		}
		catched = true
	}).Catch(TestPanic2{}, func(e E) {
		t.Error("This should not be catched")
	}).Finally(func() {
		finalized = true
	})

	if !catched {
		t.Error("Not Catched!")
	}
	if !finalized {
		t.Error("Not Finalized!")
	}

	Try(func() {
		panic(TestPanic1{"PANIC"})
	}).Catch(TestPanic1{}, func(e E) {
		if e.(error).Error() != "PANIC" {
			t.Errorf("Expected PANIC, got %s", e.(error).Error())
		}
		catched = true
	}).Catch(TestPanic2{}, func(e E) {
		t.Error("This should be skipped")
	}).Catch(TestPanic{}, func(e E) {
		t.Error("This should not be catched")
	}).Finally(nil)

	// Without finalization
	Try(func() {
		panic(TestPanic1{"PANIC"})
	}).Catch(TestPanic2{}, func(e E) {
		t.Error("This should not be catched")
	}).Catch(TestPanic{}, func(e E) {
		t.Error("This should not be catched")
	}).Catch(TestPanic1{}, func(e E) {
		if e.(error).Error() != "PANIC" {
			t.Errorf("Expected PANIC, got %s", e.(error).Error())
		}
		catched = true
	}).Finally(nil)

	if !catched {
		t.Error("Not Catched!")
	}

	finalized = false

	Try(func() {
		panic(TestPanic1{"PANIC"})
	}).Catch(TestPanic2{}, func(e E) {
		t.Error("This should not be catched")
	}).Catch(TestPanic{}, func(e E) {
		t.Error("This should not be catched")
	}).CatchAny(func(e E) {
		if e.(TestPanic1).e != "PANIC" {
			t.Error("Should receive error")
		}
	}).Finally(func() {
		finalized = true
	})

	if !finalized {
		t.Error("Should be finalized anyway")
	}
}

func TestMissingHandler(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Should recover panic")
		}
	}()

	finalized := false

	Try(func() {
		panic(TestPanic1{"PANIC"})
	}).Catch(TestPanic2{}, func(e E) {
		t.Error("This should not be catched")
	}).Catch(TestPanic{}, func(e E) {
		t.Error("This should not be catched")
	}).Finally(func() {
		finalized = true
	})

	if !finalized {
		t.Error("Should be finalized anyway")
	}

}
