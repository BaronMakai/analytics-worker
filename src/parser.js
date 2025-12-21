const { logger } = require('./logger');
const { Event } = require('./models/Event');

class Parser {
  /**
   * @param {string} eventType
   * @param {object} eventData
   * @returns {Event}
   */
  parseEvent(eventType, eventData) {
    try {
      const event = new Event(eventType, eventData);
      return event;
    } catch (error) {
      logger.error('Error parsing event:', error);
      throw error;
    }
  }

  /**
   * @param {string} eventData
   * @returns {object}
   */
  parseEventData(eventData) {
    try {
      return JSON.parse(eventData);
    } catch (error) {
      logger.error('Error parsing event data:', error);
      throw error;
    }
  }
}

module.exports = { Parser };