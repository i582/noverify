package checkers

import (
	"testing"

	"github.com/VKCOM/noverify/src/linttest"
)

func TestReturnTypes(t *testing.T) {
	test := linttest.NewSuite(t)
	test.AddFile(`<?php
/**
 * @return bool
 */
function some() { // Error
	return 10;
}

/**
 * @return string
 */
function some2() { // Error
	if (1) {
		return "";
	}
	return 10;
}

class Foo {}
class Boo {}

/**
 * @return Foo|int
 */
function some3() { // Ok
	if (1) {
		return new Foo;
	}
	return 10;
}

/**
 * @return Foo
 */
function some4() { // Error
	if (1) {
		return new Foo;
	}
	return new Boo;
}

/**
 * @return Foo
 */
function some5() { // Error
	if (1) {
		return "hello";
	}
	return some4();
}

class Test {
	/**
	 * @return self|int
	 */
	public function method1() {
		return $this;
	}

	/**
	 * @return self|int
	 */
	public function method2() {
		return $this->method1();
	}
}

interface FooIface {
	/** @return int */
    public function method();
}

class Foo implements FooIface {
	/** @return int */
	public function method() {}
}

class Coo extends Foo {
	/** @return int */
	public function method() {}
}

class Boo {
    /** @return FooIface */
    public function f(): FooIface {
        return new Foo;
    }

    /** @return FooIface[] */
    public function g(): array {
        return [new Foo];
    }

	/** @return FooIface[][] */
    public function h(): array {
        return [[new Foo]];
    }

	/** @return Foo[][] */
    public function j(): array {
        return [[new Coo]];
    }
}

`)
	test.Expect = []string{
		`The actual return types for function \some are not the same as those in PHPDoc (actual: int, PHPDoc: bool)`,
		`The actual return types for function \some2 are not the same as those in PHPDoc (actual: int|string, PHPDoc: string)`,
		`The actual return types for function \some4 are not the same as those in PHPDoc (actual: \Boo|\Foo, PHPDoc: \Foo)`,
		`The actual return types for function \some5 are not the same as those in PHPDoc (actual: \Boo|\Foo|string, PHPDoc: \Foo)`,
		`The actual return types for function method1 are not the same as those in PHPDoc (actual: \Test, PHPDoc: \Test|int)`,
	}
	test.RunAndMatch()
}

func TestReturnTypes1(t *testing.T) {
	test := linttest.NewSuite(t)
	test.AddFile(`<?php

class ChildClassOne {
  /** Some */
  public function method() {}
}

class ChildClassTwo {
  /** Some */
  public function method() {}
}

class SomeDataClass {
  /**
   * @param int $data
   * @return ChildClassOne|ChildClassTwo
   */
  public static function method(int $data) {
    if ($data) {
      return new ChildClassOne;
    }

    return new ChildClassTwo;
  }
}

class SomeClass {
  /**
   * @return ChildClassOne|null
   */
  public function method2() {
    $data = SomeDataClass::method(10);

    if ($data instanceof ChildClassOne) {
      return $data;
    }

    return null;
  }
}

`)
	test.Expect = []string{}
	test.RunAndMatch()
}
