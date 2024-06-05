import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus } from '@fortawesome/free-solid-svg-icons';
import './FloatingButton.css';

const FloatingButton = ({ onClick }) => {
return (
    <div className="floating-button" onClick = {onClick} >
      <FontAwesomeIcon icon={faPlus} />
    </div>
  );
}

export default FloatingButton

