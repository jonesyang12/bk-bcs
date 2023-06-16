!(function () {
    var svgCode = '<svg xmlns="http://www.w3.org/2000/svg" data-name="bk-bcs-color" xmlns:xlink="http://www.w3.org/1999/xlink" style="position:absolute;width:0;height:0;visibility:hidden"><symbol id="bcs-icon-color-vcluster" viewBox="0 0 1024 1024"><path fill="#F27405" fill-rule="evenodd" d="M637.833 606.815l-102.72-180.428L727.286 85.333l203.862.93-293.315 520.552zM501.57 322.37l103.43-189.63H215.636c-94.114 0-153.345 104.863-107.109 189.63h393.044zm-88.744 94.815l-203.861.927 293.32 520.555 102.715-180.429-192.174-341.053z"/></symbol><symbol id="bcs-icon-color-tencentcloud" viewBox="0 0 1024 1024"><path fill="#006FFF" d="M523.348 160c124.668 0 230.854 81.98 271.068 196.744-5.062-.339-10.153-.477-15.276-.477-23.546 0-46.432 2.924-68.34 8.441-33.793-70.7-104.568-119.375-186.404-119.375-98.1 0-180.304 69.943-201.98 163.955-24.89-3.47-58.48-7.292-83.12-1.88C262.753 266.878 381.003 160 523.349 160z"/><path fill="#00C6DA" d="M270.7 403.889c41.277 0 78.845 7.339 106.955 25.762 39.747 26.048 78.988 61.967 117.362 96.822l-60.41 59.886c-48.933-51.109-105.786-98-163.906-98-69.477 0-125.8 59.615-125.8 130.308 0 70.692 56.323 129.53 125.8 129.53 13.997 0 36.933.405 64 .66l-2.063 3.138c-12.45 18.713-37.676 46.093-75.681 82.138-119.638 0-195.922-97.646-195.922-215.466 0-117.821 93.87-214.778 209.666-214.778z"/><path fill="#01A1FF" d="M775.995 401.067c115.795 0 209.666 96.945 209.666 216.533s-93.87 216.533-209.666 216.533H245.961C483.967 596.05 621.55 464.035 658.708 438.09c33.475-23.374 73.838-37.023 117.287-37.023zM773.9 486.4c-27.405 0-52.797 8.77-73.612 23.702-21.89 15.704-105.266 95.356-250.126 238.957h323.738c70.635 0 127.896-60.655 127.896-132.526 0-71.87-57.261-130.133-127.896-130.133z"/></g></symbol><symbol id="bcs-icon-color-publiccloud" viewBox="0 0 1024 1024"><path fill="#3762B8" fill-rule="evenodd" d="M731.166 748.979v118.378c0 51.468-40.147 92.643-90.33 92.643h-269.32c-50.185 0-90.332-41.175-90.332-92.643V748.979h449.982zm-327.868 46.322c-15.055 0-28.437 12.009-28.437 29.165 0 15.44 11.71 29.166 28.437 29.166 15.055 0 28.438-13.725 28.438-29.166 0-15.44-11.71-29.165-28.438-29.165zm27.6-729.31C562.586 81.455 674.522 177.666 715.674 306.52h27.984c105.35 0 190.948 60.132 218.932 154.625 29.63 97.93-4.938 204.449-85.597 264.58-4.939 3.437-11.523 5.155-18.107 5.155-9.877 0-19.754-5.154-26.338-13.745-9.877-15.462-6.585-36.079 8.23-46.387 57.614-42.951 82.306-118.546 60.906-188.986-19.753-65.286-80.659-106.52-158.026-106.52h-49.383c-14.815 0-27.984-10.308-31.276-25.77-27.984-116.828-123.458-202.73-237.04-216.475-113.58-13.744-225.516 48.106-278.191 154.625-52.676 106.52-37.86 237.091 37.86 326.43 11.523 13.745 9.877 36.08-3.292 48.106-13.169 12.026-34.568 10.308-46.091-3.436C42.417 548.765 24.31 387.268 88.508 256.696c65.844-130.572 202.47-206.167 342.39-190.704zm211.61 403.342c50.184 0 90.331 41.175 90.331 92.644v118.377H281.184V561.977c0-51.469 40.147-92.644 90.331-92.644zm-239.21 106.369c-16.728 0-30.11 13.725-28.437 29.165 0 15.44 11.71 29.166 28.437 29.166 15.055 0 28.438-12.01 28.438-29.166 0-15.44-11.71-29.165-28.438-29.165z"/></symbol><symbol id="bcs-icon-color-k8s" viewBox="0 0 1024 1024"><path fill="#326DE6" d="M510.575 85.333c7.864 0 16.711 1.01 24.576 5.048l299.822 147.387c14.745 8.076 26.542 21.2 31.457 37.351l74.71 330.107c1.966 17.161-1.966 34.323-11.797 48.456l-207.418 263.48c-10.813 15.142-27.524 23.218-45.219 21.199H346.41c-16.711-1.01-33.423-9.086-45.22-21.2L93.774 653.681c-9.83-14.132-13.763-31.294-10.814-48.455L156.686 273.1c3.932-17.161 14.746-30.285 29.491-37.351L485.999 87.352c7.865-2.019 16.712-2.019 24.576-2.019zm-14.745 103.35c-3.367 3.405-5.33 7.521-5.892 11.755l-.168 2.55v5.11c0 4.087 1.01 8.174 2.02 12.261 1.01 2.044 1.01 4.088 1.01 6.132v1.021c2.02 11.24 2.02 22.481 2.02 34.743-1.01 3.066-2.02 6.131-5.05 8.175v2.044l-1.01 8.175c-11.109 1.022-22.218 3.065-34.337 5.109-47.466 10.218-90.893 35.765-124.22 71.53l-6.06-4.088h-1.01c-1.01 0-2.02 1.022-4.04 1.022s-4.039-1.022-6.059-2.044c-7.271-5.722-14.543-12.098-20.78-18.606l-4.468-4.896c-1.01-2.044-3.03-3.066-4.04-4.088-3.03-3.065-5.049-6.13-8.079-9.196-1.01-1.022-2.02-1.022-3.03-2.044l-1.01-1.022c-4.04-3.066-9.089-5.11-14.138-5.11-6.06 0-11.11 2.044-14.14 7.154-5.554 8.43-4.32 19.436 2.148 26.722l1.892 1.89c1.01 0 1.01 1.021 1.01 1.021s2.02 2.044 3.03 2.044c3.03 2.044 7.07 4.088 11.11 6.131l6.059 3.066c10.099 6.13 20.198 12.262 28.278 20.437 1.683 1.703 3.366 4.825 3.296 7.593l-.122.726 5.915 5.987c-1.01 2.043-2.02 3.065-3.03 5.109-31.307 50.07-44.436 109.338-35.347 167.584l-8.08 2.043c0 1.022-1.01 1.022-1.01 1.022-1.01 3.066-4.04 5.11-7.069 7.153-8.887 2.452-17.774 4.25-27.179 5.395l-7.158.736c-2.02 0-4.04 0-6.06 1.022-4.04 0-8.079 1.022-12.119 1.022-1.01 0-2.02 1.022-4.04 1.022-.757 0-.946 0-1.42.43l-.6.592c-10.315 1.897-17.147 10.844-16.453 21.112l.295 2.39c2.02 9.197 11.11 15.328 20.199 14.306 1.514 0 2.461 0 3.692-.431l1.357-.59c1.01 0 1.01 0 1.01-1.023 0-1.022 3.03 0 4.04 0 4.04-1.022 8.08-3.065 11.109-4.087 2.02-1.022 4.04-2.044 6.06-2.044h1.01c11.108-4.087 21.208-7.153 33.327-9.196h1.01c3.03 0 6.059 1.021 8.079 3.065.757 0 .947.575.994.862l.016.16 9.09-1.022c15.148 47.005 43.426 88.901 82.812 119.557 9.09 7.153 17.17 13.284 27.268 18.393l-5.05 7.153c0 1.022 1.01 1.022 1.01 1.022 2.02 3.065 2.02 7.153 1.01 10.218-4.04 10.219-10.099 20.437-16.158 29.634v1.022c-1.01 2.044-2.02 3.065-4.04 5.11-2.02 2.043-4.04 6.13-7.07 10.218-1.01 1.021-1.01 2.043-2.02 3.065 0 0 0 .575-.425.862l-.584.16c-5.05 10.218-1.01 22.48 8.08 27.59 2.02 1.022 5.049 2.044 7.069 2.044 7.181 0 13.565-4.037 17.732-9.958l1.456-2.305s0-1.021 1.01-1.021c0-1.022 1.01-2.044 2.02-3.066 1.01-4.087 3.03-7.153 4.04-11.24l2.02-6.131c3.03-11.24 8.079-21.46 13.128-31.678 2.02-3.065 5.05-5.109 8.08-6.13.757 0 .947 0 .994-.432l.016-.59 4.04-8.176c28.277 11.24 57.565 16.35 87.862 16.35 18.179 0 36.357-2.044 54.536-7.153 11.11-2.044 22.218-6.131 32.317-9.197l4.04 7.153c1.01 0 1.01 0 1.01 1.022 3.03 1.022 6.06 3.066 8.08 6.131 5.049 10.219 10.099 20.437 13.128 31.678v1.022l2.02 6.13c1.01 4.088 2.02 8.175 4.04 11.241 1.01 1.022 1.01 2.044 2.02 3.066 0 0 0 .574.426.862l.584.16c4.04 7.152 11.109 12.262 19.188 12.262 3.03 0 5.05-1.022 8.08-2.044 4.04-2.044 8.079-6.131 9.089-11.24.808-4.088.97-8.175-.033-12.263l-.977-3.065c0-1.022-1.01-1.022-1.01-1.022 0-1.022-1.01-2.044-2.02-3.066-2.02-4.087-4.04-7.153-7.07-10.218-1.01-2.044-2.02-3.066-4.04-5.11v-2.043c-7.069-9.197-12.118-19.415-16.158-29.634-1.01-3.065-1.01-7.153 1.01-10.218 0-.767.568-.958.852-1.006l.158-.016-3.03-8.175c51.506-31.677 90.893-80.726 109.072-138.972l8.079 1.022c1.01 0 1.01-1.022 1.01-1.022 2.02-2.043 5.05-3.065 8.08-3.065h1.01c11.108 2.044 22.217 5.11 32.316 9.197h1.01c2.02 1.021 4.04 2.043 6.06 2.043 4.04 2.044 7.07 4.088 11.109 5.11 1.01 0 2.02 1.021 4.04 1.021.757 0 .946 0 1.42.431l.6.591c2.02 1.022 3.03 1.022 5.05 1.022 9.088 0 17.168-6.131 20.198-14.306-.926-10.304-7.792-18.031-15.931-20.822l-2.248-.637c-1.01 0-2.02 0-2.02-1.022 0-1.021-2.02-1.021-4.04-1.021-4.04-1.022-8.079-1.022-12.119-1.022-2.02 0-4.04 0-6.06-1.022h-1.009c-11.11-1.022-23.228-3.066-34.337-6.131-2.525-.852-5.05-3.123-6.406-5.63l-.664-1.523-8.08-2.044c4.04-29.634 2.02-60.29-4.039-89.923-7.07-29.633-19.188-58.245-35.347-83.792l6.06-6.13v-1.023c0-3.065 1.01-7.153 3.03-9.196 7.27-6.54 14.542-11.772 22.33-16.742l5.947-3.695 6.06-3.066c4.04-2.043 7.07-4.087 11.109-6.13 1.01-1.023 2.02-1.023 3.03-2.045.757-.766.378-.958.142-1.437l-.143-.606c9.09-7.153 11.11-19.415 4.04-28.612-3.03-4.088-9.09-7.153-14.139-7.153-4.04 0-8.08 1.308-11.602 3.4l-2.537 1.71-1.01 1.021c-1.01 1.022-2.02 2.044-3.03 2.044-3.029 3.065-6.059 6.13-8.079 9.196-.757 1.533-2.083 2.491-3.124 3.305l-.915.783c-7.07 8.175-16.159 16.35-25.248 22.48-2.02 1.022-4.04 2.044-6.06 2.044-1.01 0-3.03 0-4.04-1.022h-1.01l-8.079 5.11c-8.08-8.175-17.168-16.35-25.248-24.525-37.367-29.633-83.823-48.027-131.29-53.136l-1.01-8.175v1.022c-3.029-2.044-4.039-5.11-5.049-8.175 0-11.24 0-22.48 2.02-34.743v-1.022c0-2.043 1.01-4.087 1.01-6.13 1.01-4.088 1.01-8.175 2.02-12.263v-6.131c1.01-10.219-7.07-20.437-17.169-21.459-6.06-1.022-12.119 2.044-17.168 7.153zm20.198 437.351c2.272 1.533 3.977 3.066 5.965 5.03l2.114 2.123 48.476 88.901c-6.059 1.022-12.119 2.044-19.188 4.088-14.139 3.065-28.278 5.109-43.427 5.109-21.208 0-43.426-4.087-63.624-10.219l49.486-90.944c5.05-6.131 13.129-8.175 20.198-4.088zm-91.903-43.94c8.08 0 16.16 5.11 17.169 13.285.808 2.452.97 5.559-.032 8.273l-.978 1.945-38.377 94.01c-35.347-23.502-63.625-59.267-77.764-100.14l99.982-17.372zm170.677-1.021l100.992 17.371c-5.05 14.306-11.11 27.59-19.189 38.83-15.149 24.525-36.357 45.984-60.595 61.312l-39.387-96.054c-2.02-8.175 2.02-16.35 9.09-19.415 3.03-1.022 6.059-2.044 9.089-2.044zm-69.685-81.748l19.189 24.524-7.07 29.634-27.268 13.284-27.267-13.284-7.07-29.634 19.189-24.524h30.297zm152.498-72.552c11.11 18.394 19.189 38.83 24.238 60.29 5.05 21.458 6.06 42.917 4.04 64.376l-95.943-27.59c-9.089-2.044-14.138-11.24-12.119-20.437 1.01-3.065 2.02-5.11 4.04-7.153l75.744-69.486zm-336.303 0l74.734 67.443c7.07 6.13 8.08 16.35 2.02 23.502-2.02 3.066-4.04 4.087-8.08 5.11l-97.962 28.611c-3.03-42.918 7.07-86.857 29.288-124.666zm143.409-94.01l-5.05 102.185c0 9.197-8.08 16.35-17.169 16.35-3.03 0-5.05-1.022-8.08-2.044l-83.822-60.29c26.258-25.545 58.575-43.939 93.922-52.114l20.199-4.087zm48.476 1.022c44.436 5.11 83.823 24.524 114.12 55.18l-82.813 59.267c-7.07 4.088-17.168 3.066-22.218-4.087-2.02-2.044-3.03-4.088-3.03-7.153l-6.06-103.207z"/></symbol><symbol id="bcs-icon-color-kubeconfig" viewBox="0 0 1024 1024"><path fill="#313238" d="M804.877 745.728l-3.446 8.82 64.971 56.588-362.944 199.513-363.562-199.812 66.808-56.29-3.342-8.094 300.096 164.954 301.42-165.679zM438.92 275.588v173.428l141.866-173.427h179.797L601.095 448.703l166.656 285.81H601.095l-92.287-188.766-69.888 76.696v112.07H303.624V275.589H438.92zm7.152-246.468v99.307L132.957 300.465l-.016 325.22-68.066 48.723-22.238 54.168-.007-477.733L446.072 29.12zm114.688-.021l403.474 221.744-.018 456.208-21.97-32.643-68.356-47.912.017-325.98-313.147-172.09V29.1z"/></symbol></svg>'
    if (document.body) {
        document.body.insertAdjacentHTML('afterbegin', svgCode)
    } else {
        document.addEventListener('DOMContentLoaded', function() {
            document.body.insertAdjacentHTML('afterbegin', svgCode)
        })
    }
})()