exports.promisify = fn => arg => {
  return new Promise((resolve, reject) => {
    fn(arg, (err, data) => {
      if (err) return reject(err);
      return resolve(data);
    });
  });
};

exports.all = promises => {
  return new Promise((resolve, reject) => {
    if (!promises?.length) return resolve(promises);

    let resolved = 0;
    const output = [];

    promises.forEach((prom, i) => {
      wrap(prom)
        .then(data => (output[i] = data))
        .catch(reject)
        .finally(() => {
          if (++resolved === promises.length) resolve(output);
        });
    });
  });
};

exports.allSettled = promises => {
  return new Promise(resolve => {
    if (!promises?.length) return resolve(promises);

    let settled = 0;
    const output = [];

    promises.forEach((prom, i) => {
      wrap(prom)
        .then(data => (output[i] = data))
        .catch(err => (output[i] = err))
        .finally(() => {
          if (++settled === promises.length) resolve(output);
        });
    });
  });
};

exports.race = promises => {
  return new Promise((resolve, reject) => {
    if (!promises?.length) return resolve(promises);

    promises.forEach(prom => {
      wrap(prom).then(resolve).catch(reject);
    });
  });
};

exports.any = promises => {
  return new Promise((resolve, reject) => {
    if (!promises?.length) return resolve(promises);

    let settled = 0;
    const output = [];

    promises.forEach((prom, i) => {
      wrap(prom)
        .then(resolve)
        .catch(err => {
          output[i] = err;
          if (++settled === promises.length) reject(output);
        });
    });
  });
};

// Convert non-promise values to promises
const wrap = prom => Promise.resolve(prom);
