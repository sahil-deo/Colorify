
# Hue Blind

This API converts normal colors into their simulated appearance for the three major types of color blindness:

- **Protanopia** (red-blind)
- **Deuteranopia** (green-blind)
- **Tritanopia** (blue-blind)

It’s built in **Go** and deployed for public use.

---

## API Usage

### Endpoint
```

GET [https://blind-hue.onrender.com/api?color=](https://blind-hue.onrender.com/api?color=)<hex>\&color=<hex>...

```

> **Note:** When using raw URLs in a browser or directly in the address bar, replace `#` with `%23` (e.g., `%23FF0000`). Most HTTP libraries in JavaScript (fetch/axios) or Python (requests) handle this encoding automatically.

### Query Parameters
- `color` — One or more hex colors (e.g. `#FF0000`).  
  You can provide multiple values by repeating the parameter.

---

### Example Request
```

GET https://blind-hue.onrender.com/api?color=%23FF0000&color=%2300FF00

````

### Example Response
```json
{
  "colors": {
    "protanopia": {
      "#FF0000": "#900000",
      "#00FF00": "#7FBF7F"
    },
    "deuteranopia": {
      "#FF0000": "#A00000",
      "#00FF00": "#7F997F"
    },
    "tritanopia": {
      "#FF0000": "#F20000",
      "#00FF00": "#6F9F9F"
    }
  }
}
````

---

## Features

* Supports multiple input colors in one request.
* Returns all three color-blindness simulations at once.
* Simple REST API, no authentication needed.

---

## References

* Brettel, Viénot and Mollon (1997). *Computerized simulation of color appearance for dichromats.*
* [Lokno’s Color Blindness Simulation Matrices (GitHub Gist)](https://gist.github.com/Lokno/df7c3bfdc9ad32558bb7)

---

## Documentation Website

Check out the docs/demo website here:
[https://blind-hue.onrender.com](https://blind-hue.onrender.com)

---

## 📜 License

MIT License — free to use in your projects.

