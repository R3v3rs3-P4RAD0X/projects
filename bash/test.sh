while read x; do
    curl "$x"
done < salt.txt