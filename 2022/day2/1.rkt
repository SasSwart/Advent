#lang racket

(define (score hand)
  (case hand
    [("A X") 4]
    [("A Y") 8]
    [("A Z") 3]
    [("B X") 1]
    [("B Y") 5]
    [("B Z") 9]
    [("C X") 7]
    [("C Y") 2]
    [("C Z") 6]
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