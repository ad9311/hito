import React from 'react';
import {useSelector} from 'react-redux';

const LandmarkDetail = () => {
  const {
    selectedLmStatus,
    selectedLm,
  } = useSelector((state) => state.landmarks);

  const detail = () => (
    <article className="landmark-detail">
      <div className="landmark-img">
        <img src={selectedLm.imgURL} alt={selectedLm.name} />
      </div>
      <div className="landmark-description">
        <h2>{selectedLm.name}</h2>
        <h3>{selectedLm.nativeName}</h3>
        <div className="wiki">
          <h4>{selectedLm.type}</h4>
          <a href={selectedLm.wikiURL}><button>Wikipedia</button></a>
        </div>
        <p>{selectedLm.description}</p>
      </div>
      <div className="detail-con">
        <p>{selectedLm.country}</p>
        <p>{selectedLm.city}</p>
      </div>
      <div className="detail-con">
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
      <div className="detail-con">
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
      <div className="landmark-actions">
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
