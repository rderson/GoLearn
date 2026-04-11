# UntisToCalendar

UntisToCalendar is a Go application that will automatically sync a WebUntis school schedule into Google Calendar for the next two weeks.

## Features

- Syncs WebUntis timetable into Google Calendar
- Creates and uses a dedicated calendar ("WebUntis")
- Updates timetable every two weeks
- Safe: never touches personal calendar events
- Works on Windows

## How it works

1. Authenticates with WebUntis using the JSON-RPC API
2. Fetches timetable data for the next 14 days
3. Converts lessons into Google Calendar events
4. Deletes outdated events in the WebUntis calendar
5. Inserts updated events

## Requirements

- Google account
- WebUntis student account
- Windows 10/11

### WebUntis credentials

You need:
- WebUntis username
- WebUntis password
- School identifier (e.g. `myschool`)

### Environment variables

Create a `.env` file in the project root:

```env
UNTIS_USER=your_username
UNTIS_PASS=your_password
UNTIS_SCHOOL=your_school
GOOGLE_CALENDAR_NAME=WebUntis
TZ=Europe/Berlin