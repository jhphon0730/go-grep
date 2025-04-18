📌 2-2. --count: 파일마다 매치된 라인 수 출력
grep -c 처럼 각 파일별로 몇 번 매치됐는지 세어줌

📌 2-3. --files-with-matches (-l)
매치된 파일 이름만 출력, 내용은 안 보여줌

📌 2-4. --exclude-dir
특정 디렉토리 (예: .git, vendor) 제외하는 옵션 유저가 직접 설정하게

📌 2-5. --context, --before-context, --after-context
매치된 줄 전후로 몇 줄 더 보여주기 (grep -C, -A, -B처럼)

📌 2-6. 파일 인코딩 처리 (고급)
나중에 한글이나 다른 인코딩 파일 읽을 때 utf-8 아닌 경우도 처리할 수 있음
