g++ day2_p2_depth.cpp -o day2_p2_depth
g++ day2_p2_horizontal.cpp -o day2_p2_horizontal
echo result:
depth=$(./day2_p2_depth) 
horizontal=$(./day2_p2_horizontal)
ans=$((depth * horizontal))
echo $ans