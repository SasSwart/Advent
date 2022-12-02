#lang racket

(define (parse-elf-calories port)
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

(define (top-n elves)
  (sort 
    elves
    > 
    #:key (lambda (calories) (calories)))
)

(define input-file (open-input-file "test.txt"))

(define elf-calories (parse-elf-calories input-file))

(define total-calories-per-elf (map (lambda (calories) (apply + (flatten calories))) elf-calories))

(displayln total-calories-per-elf)

(displayln (apply + (flatten (take (top-n elf-calories) 1))))

(displayln (apply + (flatten (take (top-n elf-calories) 3))))