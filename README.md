# photoswipestory

Static site of photos with text then swipe photos at the end. Built with [s12chung/gostatic](https://github.com/s12chung/gostatic). See demo at: http://photoswipestory.stevenchung.ca.

## Background

I made this as a birthday gift, a photo book type project, so it's not my best code.

The idea behind it is to write stories about personal photos, then show a variation of those photos at the end as a "surprise". In my case, I chose photos of personal locations and selfies at those locations at the end. But you can do themes like:

- Night vs Day (in demo)
- With or without people
- Past vs present
- In-construction vs finished

The real data is personal, so I `.gitignored` `content/markdowns`. I made a quick demo mode instead, which is set in [settings.json](settings.example.json).

## Development

To run in development, make a copy of all files with `.example` in the filename and remove `.example` in their name. Then:

```bash
make docker-install
make docker
```


Please see [s12chung/gostatic](https://github.com/s12chung/gostatic) for more information.
