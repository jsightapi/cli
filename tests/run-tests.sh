rm -f /tmp/*

echo "Running tests..."

#make file /etc/machine-id if not exists for docker conteiners and fix bug with no docker support
touch /etc/machine-id 2> /dev/null

jsight -s doc html test01.jst > /tmp/jdoc01.html
if cmp --silent /tmp/jdoc01.html ./expected/jdoc01.html; then
    echo "> Test 01: [OK]"
else
    echo "> Test 01: [FAILED] \`jsight doc html test01.jst\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/jdoc01.html`\n\n- EXPECTED:\n`head -n 5 ./expected/jdoc01.html`"
fi

jsight -s doc html test02.jst 2> /tmp/jdoc02-error
if cmp --silent /tmp/jdoc02-error ./expected/jdoc02-error; then
    echo "> Test 02: [OK]"
else
    echo "> Test 02: [FAILED] \`jsight doc html test02.jst\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/jdoc02-error`\n\n- EXPECTED:\n`cat ./expected/jdoc02-error`"
fi

jsight -s doc html test03_absent.jst 2> /tmp/jdoc03-error
if cmp --silent /tmp/jdoc03-error ./expected/jdoc03-error; then
    echo "> Test 03: [OK]"
else
    echo "> Test 03: [FAILED] \`jsight doc html test03_absent.jst\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/jdoc03-error`\n\n- EXPECTED:\n`cat ./expected/jdoc03-error`"
fi

jsight -s doc html test04.jst 2> /tmp/jdoc04-error
if cmp --silent /tmp/jdoc04-error ./expected/jdoc04-error; then
    echo "> Test 04: [OK]"
else
    echo "> Test 04: [FAILED] \`jsight doc html test04_absent.jst\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/jdoc04-error`\n\n- EXPECTED:\n`cat ./expected/jdoc04-error`"
fi

jsight -s version > /tmp/version
if cmp --silent /tmp/version ./expected/version; then
    echo "> Test 05: [OK]"
else
    echo "> Test 05: [FAILED] \`jsight version\` output does not match the expected value!\n\n- ACTUAL:\n`cat /tmp/version`\n\n- EXPECTED:\n`cat ./expected/version`"
fi


jsight -s convert openapi json test06.jst > /tmp/openapi06.json
if cmp --silent /tmp/openapi06.json ./expected/openapi06.json; then
    echo "> Test 06: [OK]"
else
    echo "> Test 06: [FAILED] \`convert openapi json test06.jst\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/openapi06.json`\n\n- EXPECTED:\n`head -n 5 ./expected/openapi06.json`"
fi

jsight -s convert openapi > /tmp/openapi07-error
if cmp --silent /tmp/openapi07-error ./expected/openapi07-error; then
    echo "> Test 07: [OK]"
else
    echo "> Test 07: [FAILED] \`convert openapi json\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/openapi07-error`\n\n- EXPECTED:\n`head -n 5 ./expected/openapi07-error`"
fi

# send statistics
jsight convert openapi json test06.jst 1> /tmp/openapi06.json 2> /dev/null
if cmp /tmp/openapi06.json ./expected/openapi06.json; then
    echo "> Test 08: [OK]"
else
    echo "> Test 08: [FAILED] \`convert openapi json test06.jst\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/openapi06.json`\n\n- EXPECTED:\n`head -n 5 ./expected/openapi06.json`"
fi

jsight convert openapi json test09.jst > /tmp/openapi09.json
if cmp --silent /tmp/openapi09.json ./expected/openapi09.json; then
    echo "> Test 09: [OK]"
else
    echo "> Test 09: [FAILED] \`convert openapi json\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/openapi09.json`\n\n- EXPECTED:\n`head -n 5 ./expected/openapi09.json`"
fi

jsight convert openapi yaml test09.jst > /tmp/openapi09.yaml
if cmp --silent /tmp/openapi09.yaml ./expected/openapi09.yaml; then
    echo "> Test 10: [OK]"
else
    echo "> Test 10: [FAILED] \`convert openapi yaml\` output does not match the expected value!\n\n- ACTUAL:\n`head -n 5 /tmp/openapi09.yaml`\n\n- EXPECTED:\n`head -n 5 ./expected/openapi09.yaml`"
fi

echo "Tests finished!"