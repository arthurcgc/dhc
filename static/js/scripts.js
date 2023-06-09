function fadeOut(el) {
    el.style.opacity = 1;
    (function fade() {
        if ((el.style.opacity -= .1) < 0) {
            el.style.display = "none";
        } else {
            requestAnimationFrame(fade);
        }
    })();
};

function fadeIn(el, display) {
    el.style.opacity = 0;
    el.style.display = display || "block";
    (function fade() {
        var val = parseFloat(el.style.opacity);
        if (!((val += .1) > 1)) {
            el.style.opacity = val;
            requestAnimationFrame(fade);
        }
    })();
};

window.addEventListener("load", () => {
    function sendData(name, email, phone) {
      const XHR = new XMLHttpRequest();

      // Define what happens on successful data submission
      XHR.addEventListener("load", (event) => {
        alert(event.target.responseText);
      });
  
      // Define what happens in case of error
      XHR.addEventListener("error", (event) => {
        alert("Oops! Something went wrong.");
      });
  
      // Set up our request
      XHR.open("POST", window.location.origin + "/email");

      XHR.setRequestHeader("Content-Type", "application/json");

      XHR.send(JSON.stringify({ "email": email, "name": name, "phone": phone }));
    }
  
    // Get the form element
    const form = document.getElementById("email-form");
  
    // Add 'submit' event handler
    form.addEventListener("submit", (event) => {
      event.preventDefault();
      var name = form.elements["name"].value
      var email = form.elements["email"].value;
      var phone = form.elements["phone"].value;

      sendData(name, email, phone);
    });
  });