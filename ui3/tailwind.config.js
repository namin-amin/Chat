/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        bgmain: "#2C363F",
        bgsecondary: "#070600",
        tirtiarry: "#d81159ff",
        secondary: "#ffbc42ff",
        primary: "#0496ffff",
        darktirtiarry: "#990c3f",
        darksecondary: "#e19100",
      },
    },
  },
  // eslint-disable-next-line no-undef
  plugins: [require("@tailwindcss/forms")],
};
