awk '{
    i=1;
    while(i <= NF){
        print $i;
        i++
    }
}' words.txt | sort | uniq -c | sort -r | awk '{print $2" "$1}'