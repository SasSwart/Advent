#lang racket

(define (calculate-elf-calories port)
  (let 
    loop (
      [line (read-line port)]
      [nth-elf (list)]
    )
    (cond 
      ; Return if EOF
      [(eof-object? line) (list)]
      ; Handle snack lines
      [(non-empty-string? line)
        (loop (read-line port)
              (cons (string->number line) nth-elf))] 
      ; Handle empty lines
      [else 
        (cons 
          nth-elf 
          (loop (read-line port)
                (list)))]
    )
  )
)

(define (calorie-ranking elves)
  (sort 
    elves
    > 
    #:key (lambda (calories) (apply + calories)))
)

(define input-file (open-input-file "test.txt"))

(define elf-calories (calculate-elf-calories input-file))

(displayln (apply + (flatten (take (calorie-ranking elf-calories) 1))))

(displayln (apply + (flatten (take (calorie-ranking elf-calories) 3))))