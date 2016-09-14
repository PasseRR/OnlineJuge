# 1st. sed
sed -n '10,10p' file.txt
# 2nd. head tail
cat file.txt | tail -n +10 | head -n 1
