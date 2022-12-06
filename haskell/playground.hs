{--
The grade function takes an Int as input and returns a String as output.
The function uses pattern matching and guards to determine the letter grade
(A, B, C, D, or F) that corresponds to the given score.
--}
grade :: Int -> String
grade score
  | score >= 90 = "A"
  | score >= 80 = "B"
  | score >= 70 = "C"
  | score >= 60 = "D"
  | otherwise = "F"

-- This function generates the list of all numbers between 1 and 100 that are divisible by either 3 or 5
divisibleBy3Or5 :: [Int]
divisibleBy3Or5 = [x | x <- [1 .. 100], x `mod` 3 == 0 || x `mod` 5 == 0]

main :: IO ()
main = do
  putStrLn "Enter your score: "
  score <- readLn
  putStrLn ("Your grade is: " ++ (grade score))

  -- Generate the list of numbers
  let numbers = divisibleBy3Or5

  -- Print the resulting list
  print numbers