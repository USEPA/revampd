//var apiUrl = <!--# echo var="API_URL" -->;

  
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
    results: []
    },
    mounted () {
      this.getUnitData("/units/findByOperatingYear",this.selected);
    },
    methods: {
      getUnitData(section, year) {
        let url = apiUrl+section+"?operatingYear="+ year;
        console.log(url);
        axios.get(url).then(response => {
        console.log(response.data);
          this.results = response.data;
        }).catch( error => { console.log(error); });
      },
    selectYear(event, year){
      
      this.getUnitData("/units/findByOperatingYear", year)
      this.showConstructedUrl("/units/findByOperatingYear",year)
    },
    showConstructedUrl(section, year){
        this.apiCall= apiUrl+section+"?operatingYear="+year;
    }
  }
  });