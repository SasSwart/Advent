#lang racket

(define input-file (open-input-file "test.txt"))

(define elf-calories
  (let 
    loop (
      [line (read-line input-file)]
      [nth-elf (list)]
    )
    (cond 
      ; Return if EOF
      [(eof-object? line) (list)]
      ; Handle snack lines
      [(non-empty-string? line)
        (loop (read-line input-file)
              (cons (string->number line) nth-elf))] 
      ; Handle empty lines
      [else 
        (cons 
          nth-elf 
          (loop (read-line input-file)
                (list)))]
    )
  )
)

(define calorie-ranking
  (sort 
    elf-calories 
    > 
    #:key (lambda (calories) (apply + calories)))
)

(displayln (apply + (flatten (take calorie-ranking 1))))

(displayln (apply + (flatten (take calorie-ranking 3))))