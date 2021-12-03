import React from 'react';

const UserForm = () => {
  return (
    <form action="/" method="post">
      <input type="hidden" name="username" value="" />
      <input type="hidden" name="csrf_token" value="" />
      <label name="name" htmlFor="name">
        <input type="text" id="name" name="name" />
      </label>
      <label name="username" htmlFor="username">
        <input type="text" id="username" name="username" />
      </label>
      <label name="password" htmlFor="password">
        <input type="password" id="password" name="password" />
      </label>
      <label name="admin" htmlFor="admin">
        <input type="checkbox" name="admin" />
      </label>
      <input type="button" value="Submit" />
    </form>
  )
};

export default UserForm;
