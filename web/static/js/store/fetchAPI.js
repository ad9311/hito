const optionalBody = (body) => {
  if (body === null) {
    return null;
  }
  return JSON.stringify(body);
};

const fetchAPI = async (method, endPoint, body = null) => {
  const request = await fetch(`http://localhost:3000/api/v1/${endPoint}`, {
    method,
    headers: {
      Accept: 'application/json; charset=utf-8',
      'Content-Type': 'application/json; charset=utf-8',
    },
    body: optionalBody(body),
  });
  return await request.json();
};

export default fetchAPI;
