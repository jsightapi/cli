rm -f /tmp/*

echo "Running tests..."

#make file /etc/machine-id if not exists for docker conteiners and fix bug with no docker support
touch /etc/machine-id 2> /dev/null

jsight -s doc html test01.jst > /tmp/test01.html
if cmp --silent /tmp/test01.html ./expected/test01.html; then
    echo "> Test 01: [OK]"
else
    echo "> Test 01: [FAILED] \`jsight doc html test01.jst\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/test01.html`\n\n- EXPECTED:\n`head -n 5 ./expected/test01.html`"
fi

jsight -s doc html test02.jst 2> /tmp/test02-error
if cmp --silent /tmp/test02-error ./expected/test02-error; then
    echo "> Test 02: [OK]"
else
    echo "> Test 02: [FAILED] \`jsight doc html test02.jst\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/test02-error`\n\n- EXPECTED:\n`cat ./expected/test02-error`"
fi

jsight -s doc html test03_absent.jst 2> /tmp/test03-error
if cmp --silent /tmp/test03-error ./expected/test03-error; then
    echo "> Test 03: [OK]"
else
    echo "> Test 03: [FAILED] \`jsight doc html test03_absent.jst\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/test03-error`\n\n- EXPECTED:\n`cat ./expected/test03-error`"
fi

jsight -s doc html test04.jst 2> /tmp/test04-error
if cmp --silent /tmp/test04-error ./expected/test04-error; then
    echo "> Test 04: [OK]"
else
    echo "> Test 04: [FAILED] \`jsight doc html test04_absent.jst\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/test04-error`\n\n- EXPECTED:\n`cat ./expected/test04-error`"
fi

jsight -s version > /tmp/version
if cmp --silent /tmp/version ./expected/version; then
    echo "> Test 05: [OK]"
else
    echo "> Test 05: [FAILED] \`jsight version\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/version`\n\n- EXPECTED:\n`cat ./expected/version`"
fi


jsight -s convert openapi json test01.jst > /tmp/test01.json
if cmp --silent /tmp/test01.json ./expected/test01.json; then
    echo "> Test 06: [OK]"
else
    echo "> Test 06: [FAILED] \`convert openapi json test01.jst\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/test01.json`\n\n- EXPECTED:\n`head -n 5 ./expected/test01.json`"
fi

jsight -s convert openapi json > /tmp/test07-error
if cmp --silent /tmp/test07-error ./expected/test07-error; then
    echo "> Test 07: [OK]"
else
    echo "> Test 07: [FAILED] \`convert openapi json\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/test07-error`\n\n- EXPECTED:\n`head -n 5 ./expected/test07-error`"
fi

# send statistics
jsight convert openapi json test01.jst 1> /tmp/test01.json 2> /dev/null
if cmp /tmp/test01.json ./expected/test01.json; then
    echo "> Test 08: [OK]"
else
    echo "> Test 08: [FAILED] \`convert openapi json test01.jst\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/test01.json`\n\n- EXPECTED:\n`head -n 5 ./expected/test01.json`"
fi

jsight convert openapi json test09.jst > /tmp/test09.json
if cmp --silent /tmp/test09.json ./expected/test09.json; then
    echo "> Test 09: [OK]"
else
    echo "> Test 09: [FAILED] \`convert openapi json\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/test09.json`\n\n- EXPECTED:\n`head -n 5 ./expected/test09.json`"
fi

echo "Tests finished!"