function Get(action, category, phone_number, callback) {
  let url = `/${BASE_PATH}/data/${action}/${category}/${phone_number}`;
  url = url.replace(/\/\//g,"/");
  // @todo: while this seems to work, in theory it replaces the double
  //        slash in the http as well.  But the browser works anyway...
  fetch(url)
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then(data => {
    callback(data);
  })
  .catch(error => {
    console.error('There was a problem with the fetch operation:', error);
  });
}
