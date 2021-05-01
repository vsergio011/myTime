import http from "../http-common";

class TutorialDataService {
  getAll() {
    const token = 'eyJhbGciOiJSUzI1NiIsImtpZCI6IjRlOWRmNWE0ZjI4YWQwMjUwNjRkNjY1NTNiY2I5YjMzOTY4NWVmOTQiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoibWluYW1lIiwicGljdHVyZSI6Imh0dHA6Ly93d3cuZXhhbXBsZS5jb20vMTIzNDU2NzgvcGhvdG8ucG5nIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL215dGltZS04MjI5MSIsImF1ZCI6Im15dGltZS04MjI5MSIsImF1dGhfdGltZSI6MTYxOTM2OTc4MiwidXNlcl9pZCI6Im5uYUdsa093OXVQSlZycWYxcHdwcHhPU1hPYzIiLCJzdWIiOiJubmFHbGtPdzl1UEpWcnFmMXB3cHB4T1NYT2MyIiwiaWF0IjoxNjE5MzY5NzgyLCJleHAiOjE2MTkzNzMzODIsImVtYWlsIjoibWllbWFpbEBnbWFsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJwaG9uZV9udW1iZXIiOiIrMzQ2NTg0NTIzNjUiLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7InBob25lIjpbIiszNDY1ODQ1MjM2NSJdLCJlbWFpbCI6WyJtaWVtYWlsQGdtYWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.SXD0gS9AqtE6ma3IEsnUKOVfqUO65ucdj5jJmqZvgOlp3OXZbGKYeVKazDCaXY2OpsENNKNH6JfKharWOhhUfCT2yAFpUpM0mc3oOpu21sn04-Qv57uDt38SpYw339ZvvXnweCc8yFocyJTH2TD04yLj01E9ibu6WilU4ENRG1_NfY7t9d9xniV5UsXWQ3VGw34bcrd2PkmSlXiSmfQILiIi032ys3jUFwRLrSN0lUJ4C9g5orImJksb3wNuijYK6oS88tMfkpnkbWeiKfwm5oyuACT_a3za4YcUvE0gjmc331QoYkK_HMNeFhwuGAD6hyUzjlv8g6DxXWEGcvS6pA'
    return http.get("/users",{ mode:'no-cors',headers: {"Authorization" : `Bearer ${token}`} });
  }

  get(id) {
    return http.get(`/user/${id}`,{
        mode:'no-cors',
        headers: {
            'Authorization': 'Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjRlOWRmNWE0ZjI4YWQwMjUwNjRkNjY1NTNiY2I5YjMzOTY4NWVmOTQiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoibWluYW1lIiwicGljdHVyZSI6Imh0dHA6Ly93d3cuZXhhbXBsZS5jb20vMTIzNDU2NzgvcGhvdG8ucG5nIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL215dGltZS04MjI5MSIsImF1ZCI6Im15dGltZS04MjI5MSIsImF1dGhfdGltZSI6MTYxOTM1NDkxNywidXNlcl9pZCI6Im5uYUdsa093OXVQSlZycWYxcHdwcHhPU1hPYzIiLCJzdWIiOiJubmFHbGtPdzl1UEpWcnFmMXB3cHB4T1NYT2MyIiwiaWF0IjoxNjE5MzU0OTE3LCJleHAiOjE2MTkzNTg1MTcsImVtYWlsIjoibWllbWFpbEBnbWFsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJwaG9uZV9udW1iZXIiOiIrMzQ2NTg0NTIzNjUiLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7InBob25lIjpbIiszNDY1ODQ1MjM2NSJdLCJlbWFpbCI6WyJtaWVtYWlsQGdtYWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.m7gx7fs5Fy313oEOfdnjXPBDxCmv4KA0A9bqaPvwow93WFxmD_AHfNyHSB0bmkIak5-doVLBJ0AImxHLKLp3lN_o9ITDd_oPaddlbsI6pxna1rrxolZj4mpjaZjT8vIyshgR8oPYXvT-6s-dNqWta2CnhGe7cDHOYCI8Hkn0pv-Yz-YL3OeFI8AqIWGkE49oYlPjoMVhLhuuN2lpVZ2PoAH7Yxy0lMI51N3y-GDfcmjbxtD2T1xBHgpoz_G0GZg_9WwcydDK6tJrTOR50QXEdLqDD8QjcGrg5yruGSrHDvKP0nAMdVtv0G-Eu1NzI86vI9RinN-Jr_VqAtrz5-u_sQ' 
          }
    });
  }

  create(data) {
    return http.post("/tutorials", data);
  }
  getkey(data) {
    return http.post("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyDZD3Dyzt5ahIfnIWWkhqYF14UJn5EoP-g", {
        email: 'miemail@gmal.com',
        lastName: 'mipass'
      });
  }

  update(id, data) {
    return http.put(`/tutorials/${id}`, data);
  }

  delete(id) {
    return http.delete(`/tutorials/${id}`);
  }

  deleteAll() {
    return http.delete(`/tutorials`);
  }

  findByTitle(title) {
    return http.get(`/tutorials?title=${title}`);
  }
}

export default new TutorialDataService();