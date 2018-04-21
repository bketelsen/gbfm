require("expose-loader?$!expose-loader?jQuery!jquery");
//require("bootstrap-sass/assets/javascripts/bootstrap.js");
//
//require("bootstrap/dist/js/bootstrap.bundle.js");
//require("mdbootstrap/js/mdb.js");
import 'bootstrap';
require("./mdb.js");

$(() => {
  activateSideNav();
});

function activateSideNav() {
  let loc = window.location;
  let path = loc.pathname;
  $("li.nav-item").removeClass("active");
  $(`.nav-item a[href='${path}']`).closest("li").addClass("active");

    //$(`#contentTabs :eq(0)`).addClass("active")
    //$(`#tabContents :eq(0)`).addClass("active")
}

