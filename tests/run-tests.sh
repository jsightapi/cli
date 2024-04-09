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
    echo "> Test 02: [FAILED] \`jsight doc html test02.jst\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/test02-error`\n\n- EXPECTED:\n`cat ./expected/test02-error`"
fi

jsight doc html test03_absent.jst 2> /tmp/test03-error
if cmp --silent /tmp/test03-error ./expected/test03-error; then
    echo "> Test 03: [OK]"
else
    echo "> Test 03: [FAILED] \`jsight doc html test03_absent.jst\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/test03-error`\n\n- EXPECTED:\n`cat ./expected/test03-error`"
fi

jsight doc html test04.jst 2> /tmp/test04-error
if cmp --silent /tmp/test04-error ./expected/test04-error; then
    echo "> Test 04: [OK]"
else
    echo "> Test 04: [FAILED] \`jsight doc html test04_absent.jst\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/test04-error`\n\n- EXPECTED:\n`cat ./expected/test04-error`"
fi

jsight version > /tmp/version
if cmp --silent /tmp/version ./expected/version; then
    echo "> Test 05: [OK]"
else
    echo "> Test 05: [FAILED] \`jsight version\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/version`\n\n- EXPECTED:\n`cat ./expected/version`"
fi




jsight convert openapi json test01.jst > /tmp/test01.json
if cmp --silent /tmp/test01.json ./expected/test01.json; then
    echo "> Test 06: [OK]"
else
    echo "> Test 06: [FAILED] \`convert openapi json test01.jst\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/test01.json`\n\n- EXPECTED:\n`head -n 5 ./expected/test01.json`"
fi


echo "Tests finished!"