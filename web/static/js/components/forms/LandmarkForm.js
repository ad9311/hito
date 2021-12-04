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
        <input
          type="text"
          id="name"
          name="name"
          placeholder="Name"
          defaultValue={onNew ? '' : selectedLm.name}
        />
      </label>
      <label name="native-name" htmlFor="native-name">
        <input
          type="text"
          id="native-name"
          name="native-name"
          placeholder="Native Name"
          defaultValue={onNew ? '' : selectedLm.nativeName}
        />
      </label>
      <label name="type" htmlFor="type">
        <input
          type="text"
          id="type"
          name="type"
          placeholder="Type"
          defaultValue={onNew ? '' : selectedLm.type}
        />
      </label>
      <label name="description" htmlFor="description">
        <textarea
          id="description"
          name="description"
          placeholder="Description"
          defaultValue={onNew ? '' : selectedLm.description}
        />
      </label>
      <label name="continent" htmlFor="continent">
        <input
          type="text"
          id="continent"
          name="continent"
          placeholder="Continent"
          defaultValue={onNew ? '' : selectedLm.continent}
        />
      </label>
      <label name="country" htmlFor="country">
        <input
          type="text"
          id="country"
          name="country"
          placeholder="Country"
          defaultValue={onNew ? '' : selectedLm.country}
        />
      </label>
      <label name="state-city" htmlFor="state-city">
        <input
          type="text"
          id="state-city"
          name="state-city"
          placeholder="State/City"
          defaultValue={onNew ? '' : selectedLm.stateCity}
        />
      </label>
      <label name="start-year" htmlFor="start-year">
        <input
          type="number"
          id="start-year"
          name="start-year"
          placeholder="Start Year"
          defaultValue={onNew ? '' : selectedLm.startYear}
        />
      </label>
      <label name="end-year" htmlFor="end-year">
        <input
          type="number"
          id="end-year"
          name="end-year"
          placeholder="End Year"
          defaultValue={onNew ? '' : selectedLm.endYear}
        />
      </label>
      <label name="area" htmlFor="area">
        <input
          type="number"
          step="any"
          id="area"
          name="area"
          placeholder="Area"
          defaultValue={onNew ? '' : selectedLm.area}
        />
      </label>
      <label name="height" htmlFor="height">
        <input
          type="number"
          step="any"
          id="height"
          name="height"
          placeholder="Height"
          defaultValue={onNew ? '' : selectedLm.height}
        />
      </label>
      <label name="wiki-url" htmlFor="wiki-url">
        <input
          type="text"
          id="wiki-url"
          name="wiki-url"
          placeholder="Wiki URL"
          defaultValue={onNew ? '' : selectedLm.wikiURL}
        />
      </label>
      <label name="img-url" htmlFor="img-url">
        <input
          type="text"
          id="img-url"
          name="img-url"
          placeholder="Image URL"
          defaultValue={onNew ? '' : selectedLm.imgURL}
        />
      </label>
      <input type="submit" value="Submit" />
    </form>
  );
};

export default LandmarkForm;
