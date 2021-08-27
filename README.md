<h1 align="center">GoWandBox</h1>

<div align="center">
  <img src="https://img.shields.io/github/languages/top/classPythonAddike/gowandbox">
  
  <a href="https://pkg.go.dev/github.com/classPythonAddike/gowandbox">
    <img src="https://pkg.go.dev/badge/github.com/classPythonAddike/gowandbox.svg" alt="Go Reference">
  </a>
  
  <a href="https://goreportcard.com/report/github.com/classPythonAddike/gowandbox">
    <img src="https://goreportcard.com/badge/github.com/classPythonAddike/gowandbox">
  </a>
  
  <a href="https://sourcegraph.com/github.com/classPythonAddike/gowandbox">
    <img src="https://sourcegraph.com/github.com/classPythonAddike/gowandbox/-/badge.svg">
  </a>
  <img src="https://www.codetriage.com/classpythonaddike/gowandbox/badges/users.svg">
  <img src="https://img.shields.io/github/license/classPythonAddike/gowandbox">
</div>

<br>

A simple wrapper for the WandBox API, written in Golang!

Documentation can be found at [classpythonaddike.github.io/gowandbox/](https://classpythonaddike.github.io/gowandbox/#/)

**Note**: This wrapper supports most of the WandBox API's endpoints, except for the `POST /permlink`. This was because, I couldn't figure out how to work with the endpoint - I kept winding up with `500 - Internal Server Errors`. I'm not sure whether it was a bug in the test data that I was using, or if its a bug in the API itself. Either way, that endpoint has not been implemented in this wrapper, and I apologise for any inconveniences. Feel free to make a PR, or open an issue if you're able to figure it out!
