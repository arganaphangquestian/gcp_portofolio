package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handler)
	appengine.Main()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w, `<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8" />
				<link rel="icon" type="image/svg+xml" href="https://storage.googleapis.com/possible-tape-350617.appspot.com/assets/favicon.17e50649.svg" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>Argana Phangquestian</title>
				<link rel="stylesheet" href="https://storage.googleapis.com/possible-tape-350617.appspot.com/assets/index.5d4a29f4.css">
			</head>
			<body>
				<div id="app" data-speed="3">
					<span class="box-1" data-speed="2"></span>
					<span class="box-2" data-speed="4"></span>
					<img src="https://storage.googleapis.com/possible-tape-350617.appspot.com/assets/profile.7e16ac13.jpg" alt="argana phangquestian" />
					<h3>Argana Phangquestian</h3>
					<p>Software Engineer</p>
					<div class="language">
						<span class="go">Go</span>
						<span class="ts">Typescript</span>
					</div>
					<div class="social">
						<a
							href="https://www.instagram.com/phangargana"
							target="_blank"
							rel="noopener noreferrer"
							title="Instagram"
						>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="#faa040"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								class="feather feather-instagram"
							>
								<rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect>
								<path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path>
								<line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line>
							</svg>
						</a>
						<a
							href="https://www.linkedin.com/in/argana-phangquestian/"
							target="_blank"
							rel="noopener noreferrer"
							title="Linkedin"
						>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="#faa040"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								class="feather feather-linkedin"
							>
								<path
									d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"
								></path>
								<rect x="2" y="9" width="4" height="12"></rect>
								<circle cx="4" cy="4" r="2"></circle>
							</svg>
						</a>
						<a
							href="https://github.com/arganaphangquestian/"
							target="_blank"
							rel="noopener noreferrer"
							title="Github"
						>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="#faa040"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								class="feather feather-github"
							>
								<path
									d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"
								></path>
							</svg>
						</a>
					</div>
				</div>
				<script>
					const els = [
						document.querySelector("#app"),
						document.querySelector(".box-1"),
						document.querySelector(".box-2"),
					];

					document.addEventListener("mousemove", (e) => {
						els.forEach((layer) => {
							const s = layer.getAttribute("data-speed");
							const x = (window.innerWidth - e.pageX * s) / 100;
							const y = (window.innerHeight - e.pageY * s) / 100;
							layer.style.transform = "translate(" + x + "px, " + y + "px)";
						});
					});
				</script>
			</body>
		</html>`)
}
