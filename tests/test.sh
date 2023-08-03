rm -f /tmp/*

echo "Running tests..."

jsight doc html test01.jst > /tmp/test01.html
if cmp --silent /tmp/test01.html ./expected/test01.html; then
    echo "> Test 01: [OK]"
else
    echo "> Test 01: [FAILED] \`jsight doc html test01.jst\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/test01.html`\n\n- EXPECTED:\n`head -n 5 ./expected/test01.html`"
fi

jsight doc html test02.jst 2> /tmp/test02-error
if cmp --silent /tmp/test02-error ./expected/test02-error; then
    echo "> Test 02: [OK]"
else
    echo "> Test 02: [FAILED] \`jsight doc html test02.jst\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/test02-error`\n\n- EXPECTED:\n`head -n 5 ./expected/test02-error`"
fi


jsight version > /tmp/version
if cmp --silent /tmp/version ./expected/version; then
    echo "> Test 03: [OK]"
else
    echo "> Test 03: [FAILED] \`jsight version\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/version`\n\n- EXPECTED:\n`cat ./expected/version`"
fi



echo "Tests finished!"