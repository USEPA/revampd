//var apiUrl = <!--# echo var="API_URL" -->;
Vue.component('paginate', VuejsPaginate)
  
  const vm = new Vue({
    el: '#app',
    data: {
    selected: '2017',
    options: [
      { text: '2017', value:'2017'},
      { text: '2018', value:'2018'},
      { text: '2019', value:'2019'}
    ],
    apiCall: '',
    results: [],
    pageNumber: 0,
    },
    mounted () {
      this.getUnitData("/units/findByOperatingYear",this.selected, limit = 20, offset=0);
    },
    methods: {

      getUnitData(section, year, limit =20, offset=0) {
        this.apiCall = apiUrl+section+"?operatingYear="+ year + "&limit="+limit+ "&offset="+offset
        console.log(this.apiCall);
        axios.get(this.apiCall).then(response => {
        console.log(response.data);
          this.results = response.data;
        }).catch( error => { console.log(error); });
      },

    selectYear(event, year){
      this.getUnitData("/units/findByOperatingYear", year)
    },

    clickCallback: function(pageNum) {
      pageOffset =  (pageNum -1) * 20

      this.getUnitData("/units/findByOperatingYear",this.selected, limit=20, offset=pageOffset)
      console.log(pageNum)
    }
 
  }
  });