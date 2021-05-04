<?php

class Foo {
}

/**
 * @return string
 */
function f(): int { // want `type in typehint and phpdoc not compatibl: cannot use string as int`
  return 2;
}

/**
 * @return string|int
 */
function f1(): int {
  return 2;
}

/**
 * @return string
 */
function f2(): Foo { // want `type in typehint and phpdoc not compatibl: cannot use string as class \Foo`
  return 2;
}

/**
 * @return Roo
 */
function f3(): Foo { // want `type in typehint and phpdoc not compatibl: cannot use class \Roo as class \Foo`
  return 2;
}

/**
 * @return int[][]
 */
function f4(): array {
  return 2;
}

/**
 * @return Roo|Boo
 */
function f5(): string { // want `type in typehint and phpdoc not compatibl: none of the possible types (\Boo|\Roo) are compatible with string`
  return 2;
}

/**
 * @return int
 */
function f6(): ?int { // want `type in typehint and phpdoc not compatibl: cannot use type int as nullable type int|null`
  return 2;
}

/**
 * @return ?int
 */
function f7(): int { // want `type in typehint and phpdoc not compatibl: cannot use type int as nullable type int|null`
  return 2;
}

/**
 * @return false
 */
function f8(): bool {
  return 2;
}

/**
 * @return true
 */
function f9(): bool {
  return 2;
}

/**
 * @return Boo|Foo|string
 */
function f10(): mixed {
  return 2;
}
