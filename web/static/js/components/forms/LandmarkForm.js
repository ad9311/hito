import React from 'react';
import {useSelector} from 'react-redux';

const LandmarkForm = () => {
  const {currentUser, csrfToken} = useSelector((state) => state.users);
  const {selectedLm} = useSelector((state) => state.landmarks);
  const {onNew} = useSelector((state) => state.forms);

  return (
    <form action="/" method="post">
      <input type="hidden" name="landmark-id" value={selectedLm.id} />
      <input type="hidden" name="csrf_token" value={csrfToken} />
      <input type="hidden" name="_method" value={onNew ? 'post' : 'patch'} />
      <input type="hidden" name="model" value="landmark" />
      <input
        type="hidden"
        name="current-user"
        value={currentUser.username}
      />
      <label name="name" htmlFor="name">
        {onNew ?
          <input
            type="text"
            id="name"
            name="name"
            placeholder="Name"
            defaultValue=""
          /> :
          <input
            type="text"
            id="name"
            name="name"
            placeholder=""
            defaultValue={selectedLm.name}
          />
        }
      </label>
      <label name="native-name" htmlFor="native-name">
        {onNew ?
          <input
            type="text"
            id="native-name"
            name="native-name"
            placeholder="Native Name"
            defaultValue=""
          /> :
          <input
            type="text"
            id="native-name"
            name="native-name"
            placeholder=""
            defaultValue={selectedLm.nativeName}
          />
        }
      </label>
      <label name="type" htmlFor="type">
        {onNew ?
          <input
            type="text"
            id="type"
            name="type"
            placeholder="Type"
            defaultValue=""
          /> :
          <input
            type="text"
            id="type"
            name="type"
            placeholder=""
            defaultValue={selectedLm.type}
          />
        }
      </label>
      <label name="description" htmlFor="description">
        {onNew ?
          <textarea
            id="description"
            name="description"
            placeholder="Description"
            defaultValue=""
          /> :
          <textarea
            id="description"
            name="description"
            placeholder=""
            defaultValue={selectedLm.description}
          />
        }
      </label>
      <label name="continent" htmlFor="continent">
        {onNew ?
          <input
            type="text"
            id="continent"
            name="continent"
            placeholder="Continent"
            defaultValue=""
          /> :
          <input
            type="text"
            id="continent"
            name="continent"
            placeholder=""
            defaultValue={selectedLm.continent}
          />
        }
      </label>
      <label name="country" htmlFor="country">
        {onNew ?
          <input
            type="text"
            id="country"
            name="country"
            placeholder="Country"
            defaultValue=""
          /> :
          <input
            type="text"
            id="country"
            name="country"
            placeholder=""
            defaultValue={selectedLm.country}
          />
        }
      </label>
      <label name="state-city" htmlFor="state-city">
        {onNew ?
          <input
            type="text"
            id="state-city"
            name="state-city"
            placeholder="State/City"
            defaultValue=""
          /> :
          <input
            type="text"
            id="state-city"
            name="state-city"
            placeholder=""
            defaultValue={selectedLm.stateCity}
          />
        }
      </label>
      <label name="start-year" htmlFor="start-year">
        {onNew ?
          <input
            type="number"
            id="start-year"
            name="start-year"
            placeholder="Start Year"
            defaultValue=""
          /> :
          <input
            type="number"
            id="start-year"
            name="start-year"
            placeholder=""
            defaultValue={selectedLm.startYear}
          />
        }
      </label>
      <label name="end-year" htmlFor="end-year">
        {onNew ?
          <input
            type="number"
            id="end-year"
            name="end-year"
            placeholder="End Year"
            defaultValue=""
          /> :
          <input
            type="number"
            id="end-year"
            name="end-year"
            placeholder=""
            defaultValue={selectedLm.endYear}
          />
        }
      </label>
      <label name="area" htmlFor="area">
        {onNew ?
          <input
            type="number"
            step="any"
            id="area"
            name="area"
            placeholder="Length"
            defaultValue=""
          /> :
          <input
            type="number"
            step="any"
            id="area"
            name="area"
            placeholder=""
            defaultValue={selectedLm.area}
          />
        }
      </label>
      <label name="height" htmlFor="height">
        {onNew ?
          <input
            type="number"
            step="any"
            id="height"
            name="height"
            placeholder="Height"
            defaultValue=""
          /> :
          <input
            type="number"
            step="any"
            id="height"
            name="height"
            placeholder=""
            defaultValue={selectedLm.height}
          />
        }
      </label>
      <label name="wiki-url" htmlFor="wiki-url">
        {onNew ?
          <input
            type="url"
            id="wiki-url"
            name="wiki-url"
            placeholder="Wiki URL"
            defaultValue=""
          /> :
          <input
            type="text"
            id="wiki-url"
            name="wiki-url"
            placeholder=""
            defaultValue={selectedLm.wikiURL}
          />
        }
      </label>
      <label name="img-url" htmlFor="img-url">
        {onNew ?
          <input
            type="url"
            id="img-url"
            name="img-url"
            placeholder="Image URL"
            defaultValue=""
          /> :
          <input
            type="text"
            id="img-url"
            name="img-url"
            placeholder=""
            defaultValue={selectedLm.imgURL}
          />
        }
      </label>
      <input type="submit" value="Submit" />
    </form>
  );
};

export default LandmarkForm;
