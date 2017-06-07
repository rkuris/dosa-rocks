# dosa-rocks service

This service is designed to replace sli.do, in a more
secure way, so people can ask anonymous questions.

It's also a perfect example of the power of DOSA.

This service assumes you're using letsencrypt to get a
SSL certificate. It also stores everything in memory,
but there's nothing to stop this from running at scale
with another supported DOSA backend (like Cassandra).

Static pages are stored in the static directory.
The API calls are handled in the API directory.

## API reference:

### POST /meeting

Creates a new meeting, expects:

```json
{
	"Handle": "(meeting handle)"
}
```
and returns
```json
{
	"MeetingID": "(meeting uuid)",
	"ManagementURL": "/meeting/(meeting uuid)/token/(delete token)"
}
```

Also sets a permanent cookie named "token" with the delete token.
Note that you can get this cookie again by visiting the ManagementURL.
With this delete token, you can delete stuff, otherwise you can only
add questions or delete your own stuff.

### POST /meeting/(meeting uuid)/question

Creates a new question, expects:
```json
{
	"Question": "(question text)"
}
```
and returns
```json
{
	"Location": "/question/meetingid/UUID"
}
```

### GET /meeting/(meeting uuid)/question

Gets a list of all questions for a meeting, ordered
by the current score and the timestamp of when they
were created.

Returns:
```json
[
	{
		"QuestionID": "(uuid)",
		"Question": "(question text)",
		"Votes": "(votes)"
	},
	...
]
```

### DELETE /meeting/(meeting uuid)/question/questionid

Removes the question indicated, requires a token
as a cookie. You can only delete questions you added
in your current session unless you have the delete
token.

### DELETE /meeting/(meeting uuid)

Removes the meeting. Meetings automatically go away
after 14 days or if we reboot the server. Requires the
delete token.
