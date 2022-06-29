export INTERVIEW_ID=$( head -n 179 streets/Buckingham_Place | tail -n 1 | sed 's/SEE INTERVIEW #//g' )
echo $INTERVIEW_ID
cat interviews/interview-$INTERVIEW_ID
echo $MAIN_SUSPECT