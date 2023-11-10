interviewi=`grep -H "licen" interviews/* | grep "\"" | cut -f1 -d ":" | cut --complement -d "-" -f1`
interview="cat interviews/interview-$interviewi"
export interview_id=$interviewi
echo $interview_id
$interview
echo $MAIN_SUSPECT
