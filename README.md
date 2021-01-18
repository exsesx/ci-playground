# Circle CI + GitHub Actions + Travis CI Demo

[![Circle CI](https://circleci.com/gh/exsesx/go-ci-example.svg?style=svg)](https://circleci.com/gh/exsesx/go-ci-example)
[![GitHub Actions](https://github.com/exsesx/go-ci-example/workflows/Go/badge.svg)](https://github.com/exsesx/go-ci-example/actions)
[![Travis CI](https://travis-ci.com/exsesx/go-ci-example.svg?branch=main)](https://travis-ci.com/exsesx/go-ci-example)

This is a golang test web application that does essentially nothing and was created for educational purposes.

I wanted to learn more about how CI works.

## CI Flow

All of them should do the same thing:

1. Checkout
2. Print Go version
3. Fetch dependencies
4. Build
5. Run tests & race detector