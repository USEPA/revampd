//var apiUrl = <!--# echo var="API_URL" -->;
function buildUrl (url) {
    return apiUrl + url + "?operatingYear=2018"
  }
  
  const vm = new Vue({
    el: '#app',
    data: {
    vueCheck: 45,
    results: []
    },
    mounted () {
      this.getPosts("/units/findByOperatingYear");
    },
    methods: {
      getPosts(section) {
        let url = API_URL+section+"?operatingYear=2018";
        console.log(url);
        axios.get(url).then(response => {
        console.log(response.data);
          this.results = response.data;
        }).catch( error => { console.log(error); });
      }
    }
  });