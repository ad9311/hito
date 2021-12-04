import React from 'react';
import {useSelector} from 'react-redux';

const UserForm = () => {
  const {currentUser, csrfToken} = useSelector((state) => state.users);
  const {onNew} = useSelector((state) => state.forms);

  return (
    <form action="/" method="post">
      <input type="hidden" name="csrf_token" value={csrfToken} />
      <input type="hidden" name="_method" value={onNew ? 'post' : 'patch'} />
      <input type="hidden" name="model" value="user" />
      <input
        type="hidden"
        name="current-user"
        value={currentUser.username}
      />
      <label name="name" htmlFor="name">
        <input
          type="text"
          id="name"
          name="name"
          placeholder="Name"
          defaultValue={onNew ? '' : currentUser.name}
        />
      </label>
      <label name="username" htmlFor="username">
        <input
          type="text"
          id="username"
          name="username"
          placeholder="Username"
          defaultValue={onNew ? '' : currentUser.username}
        />
      </label>
      <label name="password" htmlFor="password">
        {onNew ?
          <div>
            <input
              type="password"
              id="password"
              name="password"
              placeholder="Password"
            />
            <input
              type="password"
              id="password-confirmation"
              name="password-confirmation"
              placeholder="Password Confirmation"
            />
          </div> :
          <div>
            <input
              type="password"
              id="current-password"
              name="current-password"
              placeholder="Current Password"
            />
            <input
              type="password"
              id="new-password"
              name="new-password"
              placeholder="New Password"
            />
            <input
              type="password"
              id="new-password-confirmation"
              name="new-password-confirmation"
              placeholder="Password Confirmation"
            />
          </div>
        }
      </label>
      {onNew &&
        <fieldset>
          <div>
            <label name="no-admin" htmlFor="no-admin">
              <input
                type="radio"
                value="false"
                name="admin"
                id="no-admin"
                checked
              />
              Standard User
            </label>
          </div>
          <div>
            <label name="admin" htmlFor="admin">
              <input
                type="radio"
                value="true"
                name="admin"
                id="admin"
              />
              Administrator User
            </label>
          </div>
        </fieldset>
      }
      <input type="submit" value="Submit" />
    </form>
  );
};

export default UserForm;
