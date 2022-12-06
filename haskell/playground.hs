grade :: Int -> String
grade score
  | score >= 90 = "A"
  | score >= 80 = "B"
  | score >= 70 = "C"
  | score >= 60 = "D"
  | otherwise = "F"

main :: IO ()
main = do
  putStrLn "Enter your score: "
  score <- readLn
  putStrLn ("Your grade is: " ++ (grade score))