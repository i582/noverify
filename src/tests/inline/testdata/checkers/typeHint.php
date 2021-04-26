<?php

class Foo {}

/**
 * @return string
 */
function f(): int { // want `type in typehint and phpdoc not compatible`
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
function f2(): Foo { // want `type in typehint and phpdoc not compatible`
  return 2;
}

/**
 * @return Roo
 */
function f3() { // want `type in typehint and phpdoc not compatible`
  return 2;
}
