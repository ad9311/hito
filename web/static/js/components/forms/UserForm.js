import React from 'react';
import {useSelector} from 'react-redux';

const UserForm = () => {
  const {currentUser, csrfToken} = useSelector((state) => state.users);
  const {onNew} = useSelector((state) => state.forms);

  return (
    <form action="/" method="post">
      <input type="hidden" name="username" value={currentUser.username} />
      <input type="hidden" name="csrf_token" value={csrfToken} />
      <label name="name" htmlFor="name">
        {onNew ?
          <input type="text" id="name" name="name" placeholder="Name" /> :
          <input
            type="text"
            id="name"
            name="name"
            defaultValue={currentUser.name}
          />
        }
      </label>
      <label name="username" htmlFor="username">
        {onNew ?
          <input
            type="text"
            id="username"
            name="username"
            placeholder="Username"
          /> :
          <input
            type="text"
            id="username"
            name="username"
            defaultValue={currentUser.username}
          />
        }
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
              type="password-confirmation"
              id="password-confirmation"
              name="password confirmation"
              placeholder="Password Confirmation"
            />
          </div> :
          <div>
            <input
              type="old-password"
              id="old-password"
              name="old-password"
              placeholder="Old Password"
            />
            <input
              type="new-password"
              id="new-password"
              name="new-password"
              placeholder="New Password"
            />
            <input
              type="password-confirmation"
              id="password-confirmation"
              name="password confirmation"
              placeholder="Password Confirmation"
            />
          </div>
        }
      </label>
      <label name="admin" htmlFor="admin">
        {onNew ?
          <input type="checkbox" name="admin" /> :
          <input type="checkbox" name="admin" checked readOnly />
        }
      </label>
      <input type="button" value="Submit" />
    </form>
  );
};

export default UserForm;
