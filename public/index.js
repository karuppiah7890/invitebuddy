var app = new Vue({
  el: '#eventDetails',
  data: {
    eventName: '',
    eventDescription: '',
    email: ''
  },
  methods: {
    sendInvite: function() {
      console.log('sending invite!');
      axios
        .post('http://localhost:3000/invite/send', {
          event: {
            name: this.eventName,
            description: this.eventDescription,
            email: this.email
          }
        })
        .then(function(response) {
          console.log(response);
          alert('invite sent! :D');
        })
        .catch(function(error) {
          console.log(error);
          alert('an error occurred while sending invite :/');
        });
    }
  }
});
