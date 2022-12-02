#lang racket

(define (score hand)
  (case hand
    [("A X") (+ 0 3)]
    [("A Y") (+ 3 1)]
    [("A Z") (+ 6 2)]
    [("B X") (+ 0 1)]
    [("B Y") (+ 3 2)]
    [("B Z") (+ 6 3)]
    [("C X") (+ 0 2)]
    [("C Y") (+ 3 3)]
    [("C Z") (+ 6 1)]
  )
)

(define (parse-rps port) 
  (let loop (
    [line (read-line port)]
    [tally 0]
    )
    (cond 
      [(eof-object? line) tally]
      [else (loop (read-line port) (+ tally (score line)))]
    )
  )
)

(define input-file (open-input-file "test.txt"))

(define total-score (parse-rps input-file))

(displayln total-score)