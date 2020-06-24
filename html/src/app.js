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
      tableSize: 10,
      results: [],
      pageNumber: 0,
      prv: "Prev",
      nxt: "Nxt"

    },
    mounted () {
      this.getUnitData("/units/findByOperatingYear",this.selected, limit = 20, offset=0);
     
    },
    methods: {


      getUnitData(section, year, limit =20, offset=0) {
        this.apiCall = API_URL+section+"?operatingYear="+ year + "&limit="+limit+ "&offset="+offset
        console.log(this.apiCall);
        axios.get(this.apiCall).then(response => {
        console.log(response.data);
        this.tableSize= this.numberOfPages(parseInt(response.data.MetaData.total))
        this.results = response.data.Units;
        }).catch( error => { console.log(error); });
      },

      numberOfPages(totalRows, pageLength=20){
       number_of_pages = Math.ceil(totalRows / pageLength )
       return number_of_pages
      },

      selectYear(event, year){
        this.getUnitData("/units/findByOperatingYear", year)
        this.numberOfPages()
      },

      clickCallback: function(pageNum) {
        pageOffset =  (pageNum -1) * 20

        this.getUnitData("/units/findByOperatingYear",this.selected, limit=20, offset=pageOffset)
        console.log(pageNum)
      }
  
    }
    });