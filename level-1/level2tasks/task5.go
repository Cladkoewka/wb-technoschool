package level2tasks

type customError struct {
  msg string
}

func (e *customError) Error() string {
  return e.msg
}

func testErr() *customError {
  // ... do something
  return nil
}

func Task5() {
  var err error
  err = testErr()
  if err != nil {
    println("error")
    return
  }
  println("ok")
}