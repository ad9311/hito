import React from 'react';
import {useSelector} from 'react-redux';

const LandmarkDetail = () => {
  const {
    selectedLmStatus,
    selectedLm,
  } = useSelector((state) => state.landmarks);

  const detail = () => (
    <article>
      <div>
        <img src={selectedLm.imgURL} alt={selectedLm.name} />
      </div>
      <div>
        <h2>{selectedLm.name}</h2>
        <h3>{selectedLm.nativeName}</h3>
        <p>{selectedLm.type}</p>
        <p>{selectedLm.description}</p>
      </div>
      <div>
        <p>{selectedLm.continent}</p>
        <p>{selectedLm.country}</p>
        <p>{selectedLm.city}</p>
      </div>
      <div>
        <table>
          <thead>
            <tr>
              <th>Latitude</th>
              <th>Longitued</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{selectedLm.latitude}</td>
              <td>{selectedLm.longitude}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div>
        <table>
          <thead>
            <tr>
              <th>Length</th>
              <th>Width</th>
              <th>Height</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{selectedLm.length}</td>
              <td>{selectedLm.width}</td>
              <td>{selectedLm.height}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div>
        <button>Edit</button>
        <button>Delete</button>
      </div>
    </article>
  );
  return (
    <div>
      {selectedLmStatus && detail()}
    </div>
  );
};

export default LandmarkDetail;
